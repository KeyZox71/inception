package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"git.keyzox.me/42_adjoly/inception/internal/env"
	"git.keyzox.me/42_adjoly/inception/internal/log"
	"golang.org/x/sys/unix"
)

func addStreamConf(confFile, stream_output_dir string) {
	confBlock := "stream {\n  include +" + stream_output_dir + "/*.conf;\n}"

	reg, err := regexp.Compile("\\s*stream\\s*\\{")
	if err != nil {
		log.Fatal(err)
	}
	content, err := os.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}
	if reg.MatchString(string(content)) == true {
		_log.Log("note", "Stream block already present in "+confFile)
		return
	}
	file, err := os.OpenFile(confFile, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(confBlock)
	_log.Log("note", "Added stream block in config")
}

func streamTemplate(template_dir, stream_suffix string) {
	stream_output_dir := env.EnvCheck("NGINX_ENVSUBST_STREAM_OUTPUT_DIR", "/etc/nginx/stream-conf.d")

	os.Mkdir(stream_output_dir, 0755)
	if unix.Access(stream_output_dir, unix.W_OK) != nil {
		_log.Log("error", "Stream template output directory not writable "+stream_output_dir)
	}

	dir, err := os.ReadDir(template_dir)

	if err != nil {
		_log.Log("error", "Error reading "+template_dir+" folder")
	}
	addStreamConf("/etc/nginx/nginx.conf", stream_output_dir)
	_log.Log("note", "Loading stream templates...")
	for _, v := range dir {
		reg, err := regexp.Compile(stream_suffix + "$")
		if err != nil {
			log.Fatal(err)
		}
		if reg.MatchString(v.Name()) == true {
			if v.IsDir() == false {
				path := filepath.Join(template_dir, v.Name())
				content, err := os.ReadFile(path)
				if err != nil {
					_log.Log("warn", "Can't read file : "+v.Name())
					continue
				}
				cmd := exec.Command("envsubst")
				cmd.Stdin = bytes.NewReader(content)
				finalFile, err := os.Create(strings.TrimSuffix(stream_output_dir+"/"+v.Name(), stream_suffix))
				if err != nil {
					log.Fatal(err)
				}
				cmd.Stdout = finalFile
			}
		}
	}
	_log.Log("note", "Stream template loaded !")
}

func subStTemplate() {
	template_dir := env.EnvCheck("NGINX_ENVSUBST_TEMPLATE_DIR", "/etc/nginx/templates")
	template_suffix := env.EnvCheck("NGINX_ENVSUBST_TEMPLATE_SUFFIX", ".template")
	stream_suffix := env.EnvCheck("NGINX_ENVSUBST_STREAM_TEMPLATE_SUFFIX", ".stream-template")
	output_dir := env.EnvCheck("NGINX_ENVSUBST_OUTPUT_DIR", "/etc/nginx/http.d")

	unix.Access(template_dir, unix.W_OK)
	unix.Access(output_dir, unix.W_OK)

	dir, err := os.ReadDir(template_dir)

	if err != nil {
		_log.Log("error", "Error reading "+template_dir+" folder")
	}
	_log.Log("note", "Loading templates...")
	for _, v := range dir {
		reg, err := regexp.Compile(template_suffix + "$")
		if err != nil {
			log.Fatal(err)
		}
		if reg.MatchString(v.Name()) == true {
			if v.IsDir() == false {
				path := filepath.Join(template_dir, v.Name())
				content, err := os.ReadFile(path)
				if err != nil {
					_log.Log("warn", "Can't read file : "+v.Name())
					continue
				}
				cmd := exec.Command("envsubst")
				cmd.Stdin = bytes.NewReader(content)
				finalFile, err := os.Create(strings.TrimSuffix(output_dir+"/"+v.Name(), template_suffix))
				if err != nil {
					log.Fatal(err)
				}
				cmd.Stdout = finalFile
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	_log.Log("note", "Template loaded !")
	i := 0
	
	reg, err := regexp.Compile(stream_suffix + "$")
	for _, v := range dir {
		if reg.MatchString(v.Name()) {
			i++;
		}
	}
	if i == 0 {
		_log.Log("note", "No stream template found, skipping...")
		return
	}
	streamTemplate(template_dir, stream_suffix)
}

func main() {
	args := os.Args

	if args[1] == "nginx" || args[1] == "nginx-debug" {
		_log.Log("note", "Entrypoint script for NGINX Server started")

		subStTemplate()

		dir, err := os.ReadDir("/docker-entrypoint.d")
		if err != nil {
			log.Fatal(err)
		}
		_log.Log("note", "Running entrypoint scripts")
		for _, v := range dir {
			os.Chmod("/docker-entrypoint.d/" + v.Name(), 0755)
			cmd := exec.Command("/docker-entrypoint.d/" + v.Name())
			cmd.Env = os.Environ()
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error running script(%s): %v\n", v.Name(), err)
				os.Exit(1)
			}
		}
		_log.Log("note", "Starting NGINX")
	}
	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running NGINX: %v\n", err)
		os.Exit(1)
	}
}

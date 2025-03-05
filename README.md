# Inception Project

This repository contains the Inception project, which requires specific environment variables and secrets to function correctly. Below are the instructions for setting up the necessary environment variables and secrets.

## Environment Variables

Create a `.env` file in the `srcs` directory of the project with the following environment variables:

```plaintext
DB_USER=your_database_user
DB_NAME=your_database_name

WP_ADMIN=your_wordpress_admin_user
WP_MAIL=your_wordpress_admin_email

FTP_USER=your_ftp_user
```

## Secrets

The secrets/ directory contains sensitive information required for various services. Ensure that the following files are present with the appropriate secrets:

### Directory Structure

```
secrets/
├── borg
│   └── passphrase.txt        # Borg backup passphrase
├── db
│   ├── root_pass.txt         # Database root password
│   └── user_pass.txt         # Database user password
├── ftp
│   └── pass.txt              # FTP password
└── wp
    └── admin_pass.txt        # WordPress admin password
```

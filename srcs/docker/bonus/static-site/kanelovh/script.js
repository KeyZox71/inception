function init()
{
    var img = document.querySelector("#kanel");
    const random = Math.floor(Math.random() * 70) + 1;
	img.src = 'https://kanel.ovh/img/'+random+'.jpg';
    const { naturalWidth, naturalHeight } = img;
    const isLandscape = naturalWidth / naturalHeight > 1;

    if (isLandscape) {
        img.style.width = window.innerWidth+'px';
        img.style.height = 'auto';
    } else {
        img.style.height = window.innerHeight+'px';
        img.style.width = 'auto';
    }

	if (img.src == 'https://kanel.ovh/img/70.jpg')
    {
        img.style.height = 'auto';
        img.style.width = '100%';
    }
}

init();

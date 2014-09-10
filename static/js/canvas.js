window.onload = function() {
	var size = 1001;
	var canvas = document.getElementById('canvas');
	var ctx = canvas.getContext('2d');
	var w = size;
	var h = size;

	canvas.width  = w;
	canvas.height = h;

	spiral = getSpiral(size);
	drawBackground(ctx, w, h);
	draw(ctx, spiral, size);
}


function getSpiral(size) {

	// Création de l'objet XmlHttpRequest
	var xhr = getXMLHttpRequest();

	// Chargement du fichier
	xhr.open("GET", '/api/ulam?size='+size, false);
	xhr.send(null);
	if(xhr.readyState != 4 || (xhr.status != 200 && xhr.status != 0)) // Code == 0 en local
		throw new Error("Impossible de charger la carte nommée \"" + nom + "\" (code HTTP : " + xhr.status + ").");
	var mapJsonData = xhr.responseText;

	// Analyse des données
	var spiralData = JSON.parse(mapJsonData);

	return spiralData;	
}

function drawBackground(ctx, w, h) {
	var imageData = ctx.createImageData(w, h)

	var index = 0;
	var pixels = w * h * 4
	while (index < pixels) {
		imageData.data[index]    = 0
		imageData.data[index +1] = 0
		imageData.data[index +2] = 0
		imageData.data[index +3] = 255		
		index += 4;
	}

	ctx.putImageData(imageData, 0, 0)
}

function draw(ctx, spiral, size) {
	var imageData = ctx.createImageData(1, 1)

	// Set red color
	imageData.data[0] = 255
	imageData.data[1] = 0
	imageData.data[2] = 0
	imageData.data[3] = 255

	ctx.putImageData(imageData, 5, 5)

	var x = 0;
	while (x < size) {

		var y = 0;
		while (y < size) {

			if (spiral[x][y] == true)
				ctx.putImageData(imageData, x, y)

			y++;
		}

		x++;
	}
}
package main

import (
	"github.com/faiface/pixel"
	"image"
	"log"
	"os"
	"path"
	"path/filepath"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

var sprites map[string]*pixel.Sprite

func loadSprites() map[string]*pixel.Sprite {
	dir := "./resources/PNG/Retina/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	sprites := make(map[string]*pixel.Sprite, len(entries))

	for _, entry := range entries {
		filename, _ := filepath.Abs(path.Join(dir, entry.Name()))
		log.Printf("Loading sprite \"%s\"\n", filename)
		sprite, err := loadSprite(filename)
		if err != nil {
			panic(err)
		}

		sprites[fileNameWithoutExtSliceNotation(entry.Name())] = sprite
	}

	log.Printf("Loaded %d sprites\n", len(sprites))

	return sprites
}

func loadSprite(path string) (*pixel.Sprite, error) {
	pic, err := loadPicture(path)
	if err != nil {
		log.Printf("Cannot load picture \"%s\"\n", err.Error())
		return nil, err
	}

	return pixel.NewSprite(pic, pic.Bounds()), nil
}

func init() {
	sprites = loadSprites()
}

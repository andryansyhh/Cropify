# Cropify

**Cropify** is a lightweight Go tool that crops images based on black border detection — using only Go's standard library.  
It’s a great example of manual pixel inspection, perfect for learning image processing in Go without any external dependencies.

---

## What It Does
Automatically detects the black rectangular border in an image  
Crops the image to include only the area inside the border (the border itself stays visible)  
Optionally logs all detected black pixel coordinates to a file

---

## How to Use
1. Clone this repo
2. Put your image as `image.png` in the root folder
3. Run:

```bash
go run main.go

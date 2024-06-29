# Waifu Fetcher

## Description

Waifu Fetcher is a simple project built with the Go programming language. This project utilizes the [waifu.pics](https://waifu.pics/) API to fetch waifu images concurrently using goroutines and channels, demonstrating the power and efficiency of Go's concurrency model.

## Features

- **Concurrency**: Fetch multiple waifu images concurrently.
- **Customizable**: Specify the number of images and the number of goroutines.
- **API Integration**: Integrates with the waifu.pics API.

## Usage

1. **Clone the repository**:

   ```bash
   git clone https://github.com/shironxn/waifu-fetcher
   ```

2. **Navigate to the project directory**:

   ```bash
   cd waifu-fetcher
   ```

3. **Run the application**:

   ```bash
   make run
   ```

4. **Follow the prompts**:
   - Enter the type (default is `sfw`).
   - Enter the category (default is `waifu`).
   - Enter the number of URLs to fetch (default is `1`).
   - Enter the number of goroutines to use (default is `1`).

## Example

Here is an example of the program in action:

```bash
Type (default=sfw): sfw
Category (default=waifu): waifu
How many URLs (default=1): 10
How many goroutines (default=1): 2

URL 1: https://i.waifu.pics/Jngw47s.png
URL 2: https://i.waifu.pics/E_U9eeg.jpg
URL 3: https://i.waifu.pics/rUfJc7w.jpg
URL 4: https://i.waifu.pics/GLGHJqM.jpg
URL 5: https://i.waifu.pics/f13ZjEw.jpg
URL 6: https://i.waifu.pics/9e8~~fj.png
URL 7: https://i.waifu.pics/Dj~3I8q.jpg
URL 8: https://i.waifu.pics/1goGv7V.jpg
URL 9: https://i.waifu.pics/-eOx7uI.jpg
URL 10: https://i.waifu.pics/P65Fb_X.jpg
```


from bs4 import BeautifulSoup
import requests as reqs


def main():
    page = reqs.get("https://kodem-tcg.com/raices-misticas")
    parser = BeautifulSoup(page.text, "html.parser")
    images = parser.find_all("img")

    for image in images:
        name = images[1].get("alt").replace(",", "").replace(" ", "_")
        print(name)
        print(images[1].get("src"))


if __name__ == "__main__":
    main()

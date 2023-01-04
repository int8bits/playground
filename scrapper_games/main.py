# from selenium import webdriver
# from selenium.webdriver.common.by import By
import re

import requests as reqs
from bs4 import BeautifulSoup


WHITELIST = [
    # "nintendo-ds",
    # "playstation-portable",
    # "gameboy-advance",
    # "gamecube",
    # "nintendo-wii",
    # "super-nintendo",
    # "playstation-2",
    # "nintendo-64",
    # "playstation",
    "nintendo",
    # "sega-genesis",
    # "gameboy-color",
    # "dreamcast",
    # "gameboy",
    # "mame-037b11",
    # "sega-saturn",
    # "atari-2600",
    # "microsoft-xbox",
    # "snk-neo-geo",
    # "amiga-500",
    # "sega-master-system",
    # "zx-spectrum",
    # "game-gear",
    # "commodore-64",
    # "turbografx-16",
    # "amstrad-cpc",
    # "capcom-play-system",
    # "nokia-n-gage",
    # "atari-800",
    # "sega-32x",
    # "colecovision",
    # "capcom-play-system-2",
    # "msx-computer",
    # "wonderswan",
    # "atari-7800-prosystem",
    # "nintendo-pokemon-mini",
    # "nintendo-famicom-disk-system",
    # "neo-geo-pocket-color",
    # "atari-lynx",
    # "msx-2",
    # "atari-jaguar",
    # "nintendo-virtual-boy",
    # "intellivision",
    # "apple-ii",
    # "atari-5200-supersystem",
    # "atari-st",
    # "commodore-vic20",
    # "sega-pico",
    # "capcom-play-system-3",
    # "bbc-micro",
    # "sega-sg1000",
    # "magnavox-odyssey-2",
    # "acorn-electron",
    # "sharp-x68000",
    # "gce-vectrex",
    # "acorn-8-bit",
    # "apple-ii-gs",
    # "acorn-archimedes",
    # "nintendo-3ds",
    # "tangerine-oric",
    # "tiger-game-com",
    # "vtech-v-smile",
    # "dragon-data-dragon",
    # "colecovision-adam",
    # "sinclair-zx81",
    # "robotron-z1013",
    # "neo-geo-pocket",
    # "thomson-mo5",
    # "miles-gordon-sam-coupe",
    # "watara-supervision",
    # "fairchild-channel-f",
    # "amstrad-gx4000",
    # "sega-visual-memory-system",
    # "philips-videopac",
    # "sufami-turbo",
    # "tandy-color-computer",
    # "z-machine-infocom",
    # "super-grafx",
    # "epoch-super-cassette-vision",
    # "bally-pro-arcade-astrocade",
    # "sharp-mz-700",
    # "emerson-arcadia-2001",
    # "commodore-plus4-c16",
    # "gamepark-gp32",
    # "vtech-creativision",
    # "pel-varazdin-orao",
    # "memotech-mtx512",
    # "camputers-lynx",
    # "elektronika-bk",
    # "commodore-pet",
    # "entex-adventure-vision",
    # "mattel-aquarius",
    # "funtech-super-acan",
    # "hartung-game-master",
    # "galaksija",
    # "interact-family-computer",
    # "casio-pv1000",
    # "apple-1",
    # "casio-loopy",
    # "sega-super-control-station",
    # "wang-vs",
    # "commodore-max-machine",
    # "luxor-abc-800",
    # "rca-studio-ii",
    # "kaypro-ii",
    # "nintendo-ds",
    # "playstation-portable",
    # "gameboy-advance",
    # "gamecube",
    # "nintendo-wii",
    # "super-nintendo",
    # "playstation-2",
    # "nintendo-64",
    # "gameboy-advance",
]
BASE_URL = "https://www.romsgames.net"
URL = f"{BASE_URL}/roms/"


def get_games_urls(url, games=None):
    print(url)
    print(games)
    if url == "#":
        pass


def get_urls_from_paginators(page_raw):
    soup = BeautifulSoup(page_raw.text, "html.parser")
    paginator = soup.find_all("ul", class_="pagination")
    games = soup.find_all("ul", class_="rg-gamelist")
    # print(page_raw.text)
    for pag in paginator:
        print(pag.get("class"))

        if "prenex" in pag.get("class"):
            continue

        pag = str(pag)
        pag_links = BeautifulSoup(pag, "html.parser")
        links = pag_links.find_all("a", href=True)

        for link in links:
            get_games_urls(link, games)


def main():
    print("using another tools")
    req = reqs.get(URL)
    soup = BeautifulSoup(req.text, "html.parser")
    job_elements = soup.find_all("a", href=True)

    for job_element in job_elements:
        is_rom_page = re.compile(r'/roms/*/')
        if is_rom_page.match(job_element["href"]):
            href = job_element["href"]
            name = href.split("/")
            name = list(filter(None, name)).pop()

            if name in WHITELIST:
                page = reqs.get(f"{BASE_URL}{href}")
                get_urls_from_paginators(page)
                break


if __name__ == "__main__":
    main()
    # driver = webdriver.Firefox(executable_path='./geckodriver')
    # driver.get("https://www.romsgames.net/roms/")
    # try:
    #     consoles = driver.find_elements(
    #         By.XPATH, "//a[contains(@href,'/roms')]"
    #     )

    #     for console in consoles:
    #         name = console.get_attribute('href').split('/')
    #         name = list(filter(None, name)).pop()
    #         print(name)
    #         if name in WHITELIST:
    #             console.click()
    #             driver.back()
    #             driver.implicitly_wait(1)
    # except Exception as e:
    #     print(e)

    # driver.quit()

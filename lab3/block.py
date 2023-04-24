from PIL import Image
from hashlib import sha1
import random


def ecb(obraz_w_bitach, rozmiar, wielkosc_blokow, klucze):
    zaszyfrowany_obraz = []

    szerokosc = rozmiar[0]
    wysokosc = rozmiar[1]

    for wiersz in range(szerokosc):
        for kolumna in range(wysokosc):
            pozycja_pixela = wiersz * szerokosc + kolumna
            oryginalny_pixel = obraz_w_bitach[pozycja_pixela]

            nowy_pixel = oryginalny_pixel ^ klucze[wiersz % wielkosc_blokow][kolumna % wielkosc_blokow]
            zaszyfrowany_obraz.append(nowy_pixel)

    obraz_ecb = Image.new("L", rozmiar)
    obraz_ecb.putdata(zaszyfrowany_obraz)
    obraz_ecb.save("ecb_crypto.bmp")
    print("Zaszyfrowano ECB")


def cbc(obraz_w_bitach, rozmiar, klucze, wielkosc_blokow):
    szerokosc = rozmiar[0]
    wysokosc = rozmiar[1]

    #poczatek_szyfru = random.randint(1, 50)
    wektor_poczatkowy = 0
    zaszyfrowany_obraz = [obraz_w_bitach[0] ^ wektor_poczatkowy]

    for i in range(szerokosc * wysokosc):
        blok = zaszyfrowany_obraz[i - 1] ^ obraz_w_bitach[i] ^ klucze[i % wielkosc_blokow**2 // wielkosc_blokow][i % wielkosc_blokow]

        zaszyfrowany_obraz.append(blok)

    obraz_cbc = Image.new("L", rozmiar)
    obraz_cbc.putdata(zaszyfrowany_obraz[1:])
    obraz_cbc.save("cbc_crypto.bmp")

    print("Zaszyfrowano CBC")


def generowanie_klucza(wielkosc_blokow):
    klucze = []

    for i in range(wielkosc_blokow):
        klucz = sha1(str(random.random()).encode("UTF-8")).digest()
        klucze.append(klucz)
        print(klucz)
    return klucze


if __name__ == "__main__":
    obraz = Image.open("images/ga.bmp")
    wielkosc_blokow = 8
    obraz_w_bitach = obraz.tobytes()
    rozmiar = obraz.size

    klucze = generowanie_klucza(wielkosc_blokow)
    ecb(obraz_w_bitach, rozmiar, wielkosc_blokow, klucze)
    cbc(obraz_w_bitach, rozmiar, klucze, wielkosc_blokow)
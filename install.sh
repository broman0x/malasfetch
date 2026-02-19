#!/bin/bash

HIJAU='\033[0;32m'
BIRU='\033[0;34m'
MERAH='\033[0;31m'
KUNING='\033[0;33m'
RESET='\033[0m'

REPO="broman0x/malasfetch"
BINARY_NAME="malasfetch"
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"

echo -e "${BIRU}--- Memulai Instalasi $BINARY_NAME ---${RESET}"

install_via_binary() {
    echo -e "${BIRU}Mencoba download binary dari GitHub Release...${RESET}"    
    LATEST_TAG=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    
    if [ -z "$LATEST_TAG" ]; then
        echo -e "${KUNING}Gagal dapet info release terbaru. Mungkin belum ada release?${RESET}"
        return 1
    fi

    echo -e "${BIRU}Dapet versi terbaru: $LATEST_TAG${RESET}"
    DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_TAG/$BINARY_NAME"
    curl -L "$DOWNLOAD_URL" -o "$BINARY_NAME"
    
    if [ $? -ne 0 ]; then
        echo -e "${KUNING}Gagal download binary.${RESET}"
        return 1
    fi

    chmod +x "$BINARY_NAME"
    echo -e "${BIRU}Memindahkan binary ke $INSTALL_PATH.${RESET}"
    sudo mv "$BINARY_NAME" "$INSTALL_PATH"
    
    return $?
}

install_via_source() {
    echo -e "${BIRU}Mencoba build dari source...${RESET}"
    
    if ! command -v go &> /dev/null; then
        echo -e "${MERAH}Duh, Go (Golang) belum ada di sistem kamu. Gak bisa build dari source, Bos!${RESET}"
        return 1
    fi

    go build -o "$BINARY_NAME" main.go
    if [ $? -ne 0 ]; then
        echo -e "${MERAH}Gagal nge-build project-nya.${RESET}"
        return 1
    fi

    echo -e "${BIRU}Memindahkan binary ke $INSTALL_PATH ${RESET}" 
    sudo mv "$BINARY_NAME" "$INSTALL_PATH"
    
    return $?
}

if install_via_binary; then
    echo -e "${HIJAU}MANTAP! $BINARY_NAME sudah terinstal via Binary Release.${RESET}"
else
    echo -e "${KUNING}Gagal instal lewat binary, coba build dari source ya...${RESET}"
    if install_via_source; then
        echo -e "${HIJAU}MANTAP! $BINARY_NAME sudah terinstal via Source Build.${RESET}"
    else
        echo -e "${MERAH}Waduh, semua cara udah dicoba tapi gagal semua. Cek koneksi atau repo-nya${RESET}"
        exit 1
    fi
fi

echo -e "${BIRU}Tinggal ketik 'malasfetch' buat pamer ke fesnuk!${RESET}"

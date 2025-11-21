# 1. Gunakan image Golang versi kecil (Alpine)
FROM golang:1.22-alpine

# 2. Set folder kerja di dalam kontainer
WORKDIR /app

# 3. Copy file dependency dulu (biar cache build lebih cepat)
COPY go.mod go.sum ./
RUN go mod download

# 4. Copy seluruh source code
COPY . .

# 5. Build aplikasi menjadi binary bernama 'main'
RUN go build -o main .

# 6. Hugging Face mengharapkan aplikasi berjalan di user non-root (opsional tapi recommended)
RUN adduser -D -u 1000 user
USER user
ENV PATH="/app:${PATH}"

# 7. Expose port 7860
EXPOSE 7860

# 8. Jalankan aplikasi
CMD ["./main"]
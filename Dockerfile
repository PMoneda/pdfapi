FROM golang:1.8

ADD wkhtmltox.tar.xz /usr/local/
ADD pdfapi /usr/local
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl \
       xz-utils \
       fontconfig libfontconfig1 libfreetype6 \
       libpng12-0 libjpeg62-turbo \
       libssl1.0.0 libx11-6 libxext6 libxrender1 \
       xfonts-base xfonts-75dpi && \            
    apt-get purge -y curl && \
    rm -rf /var/lib/apt/lists/*
ENV PATH $PATH:/usr/local/wkhtmltox/bin
ENV PATH $PATH:/usr/local/

ENTRYPOINT ["pdfapi"]
EXPOSE 8080

# pdfapi
API to converts HTML to PDF using wkhtml2pdf

#### run
  ```dockerfile
  docker run -d -p 8080:8080 pmoneda/pdfapi
  ```

Test:
  ```curl
  curl -X POST \
http://localhost:8080/topdf \
-d '<html>
<body>
<p><b>Get My PDF</b></p>
</body>
</html>'

  ```

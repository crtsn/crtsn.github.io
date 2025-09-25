#!/usr/bin/env python3
import http.server

PORT = 8080

if __name__ == "__main__":
    try:
        with http.server.ThreadingHTTPServer(("", PORT), http.server.SimpleHTTPRequestHandler) as httpd:
            with open("404.html") as html_file:
                httpd.RequestHandlerClass.error_message_format = html_file.read()
            print(f"Serving at http://127.0.0.1:{PORT}/ (Ctrl+C to stop)")
            httpd.serve_forever()
    except KeyboardInterrupt:
        print("Keyboard interrupt detected")

# dnsq

Web app to make various DNS queries and get back JSON formatted responses. Handy for doing DNS lookups in multi-language ecosystems without having to re-implement DNS lookups in each application. Just query the app, get back an answer, and parse it.

## Usage

Works with `POST` or `GET` verbs. Always with `q` as the input.

    curl "http://localhost:8080/mx?q=gmail.com"
    > {"answer":["gmail-smtp-in.l.google.com","alt1.gmail-smtp-in.l.google.com","alt2.gmail-smtp-in.l.google.com","alt3.gmail-smtp-in.l.google.com","alt4.gmail-smtp-in.l.google.com"]}

    curl "http://localhost:8080/cname?q=research.swtch.com"
    > {"answer":"ghs.google.com"}

    curl "http://localhost:8080/ip?q=patricktulskie.com"
    > {"answer":["104.21.78.121","172.67.220.247","2606:4700:3034::ac43:dcf7","2606:4700:3035::6815:4e79"]}

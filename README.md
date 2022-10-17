# pf2e-rmen.exe
pf2e-rmen.exe entfernt die englischen Überbleibsel aus der deutschen Übersetzung.
Beispiel:
* Aus "Tanzende Lichter/Dancing Lights" wird "Tanzende Lichter"
* Aus "Langschwert/Longsword" wird "Langschert"
* Aus "Druide/Druid" wird "Druide"
* etc....

Das Ausführen der exe verändert den Inhalt der json-Dateien des lang-de-pf2e-Moduls. Dieser Vergang ist somit **destruktiv**. Nach jedem Update des Übersetzungsmoduls müsst ihr die pf2e-rmen.exe erneut ausführen.
Die pf2e-rmen.exe ist nur mit **Foundry V10** kompatibel.

DAS VERWENDEN DER DATEIEN VON DIESEM REPO GESCHIEHT AUF EIGENE GEFAHR.

# Wo ist die pf2e-rmen.exe?
Ihr findet die Datei im cmd/pf2e-rmen Verzeichnis.

# Was ist noch zu tun?
Ihr müsst die pf2e-ref.config.json Datei anpassen und "DataFolder" auf euer Data-Verzeichnis anpassen. Danach seid ihr bereit die EXE auszuführen.

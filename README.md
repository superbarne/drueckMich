# Einleitung
drückMich ist eine Webanwendung zur Speicherung und Verwaltung von Web-Lesezeichen. Bei Betätigung eines sog. Browser-Action-Buttons wird die URL der aktuell im Browser angezeigten Seite clientseitig erfasst. Diese URL wird an eine Serveranwendung gesendet, um dort weiter analysiert und in einer Datenbank gespeichert zu werden. Gleichzeitig wird die drückMich Anwendung in einem neuen Browser-Tab geöffnet - falls nicht bereits geschehen. 

Neben der URL speichert die Anwendung zusätzliche Informationen über eine Seite, etwa:

 - eine Kurzbeschreibung 
 - ein Icon- oder Logo-Image
 - Kategorien
 - aus Bildern einer Seite extrahierte Daten, z.B. GPS-Positionen 

Die drückMich Anwendung listet die gespeicherten Lesezeichen auf und erlaubt, die entsprechenden URLs später erneut aufzurufen. Hierbei werden verschiedene Auflistungen angeboten: sortiert, in Kategorien, im Abstand zu einer Position, usw. 

Desweiteren ist es möglich, den drückMich-Datenbestand zu durchsuchen (URLs Kurzbeschreibungen, Kategorien, Umkreissuche, ... ).

Benutzer haben nur Zugriff auf die selbst erstellten Lesezeichen. drückMich stellt daher ein eigenes Registrierungs- und Anmeldesystem bereit. 

## Allgemeine Anforderungen

Folgende Architektur soll umgesetzt werden:

 - Das drückMich Gesamtsystem soll aus den Komponenten Client, Server und Datenbank bestehen.
 - Der Client soll aus einer Rich-Client Anwendung (HTML, CSS, JS) und einer Browser-Extension bestehen. 
 - Der drückMich-Server soll einen HTTP-Server bereitstellen, mit der Datenbank kommunizieren und als HTTP-Client Anfragen an externe Server stellen können

# Funktionsanforderungen 
## Benutzerverwaltung

 - Eine Benutzerverwaltung soll sicherstellen, dass nur registrierte Benutzer*innen Zugang zu drückMich erhalten. 
 -  Neue Benutzer*innen sollen sich durch Angabe eines Benutzernamens und eines Kennworts registrieren können. 
 - Für die Namens- und Kennwortfestlegung soll ein sinnvoller Zeichenvorrat und eine sinnvolle Zeichenkettenlänge definiert werden. Bei der Registrierung soll dies angezeigt und überprüft werden. 
 - Eine Verschlüsselung der Kennwörter, Benutzerdaten oder sonstiger Daten soll ausdrücklich nicht erfolgen.
 - Angemeldete Benutzer*innen sollen nur zu ihren eigenen Lesezeichen und den damit gespeicherten Informationen Zugang bekommen.
 - Benutzer*innen sollen ihr Konto mit allen damit verbundenen Dateninhalten löschen können.

## Browser Extension

 - Es soll eine Browser-Extension entwickelt werden, die als sog. entpackte Erweiterung im Entwicklermodus des Browsers von einem lokalen Datenträger installiert werden kann.
 - Diese Extension soll allein über einen Action-Button bedient werden.
 - Die Extension soll ohne Popup Fenster auskommen. 
 - Bei Betätigung des Action-Buttons soll die Extension die URL der im aktiven Tab ausgeführten Seite ermitteln und an den drückMich-Server senden.
 - Die Extension soll nur die URL inkl. einer Benutzerkennung an den drückMich-Server senden, d.h. der Benutzer muss dort angemeldet sein. 
 - Bei Betätigung des Action-Buttons soll auch der drückMich-Rich-Client in einem Tab starten, falls dies nicht bereits der Fall ist. 

## Rich-Client

 - Der drückMich-Rich-Client soll das User-Interface zur Bedienung der drückMich-Server Anwendung bilden. 
 - Der Rich-Client soll alle von der Benutzer*in gespeicherten Lesezeichen als Links anzeigen können, sodass mit einem Klick die entsprechende Seite in einem neuen Tab geöffnet wird.
 - Lesezeichen sollen mit einem Icon, einem Titel, der URL und einer Kurzbeschreibung dargestellt werden. 
 - Lesezeichenlisten sollen unterschiedlich dargestellt werden können: 
   - alphabetisch sortiert nach Titeln
   - nur Lesezeichen bestimmter Kategorien
   - sortiert nach Alter (Erstelldatum der Lesezeichen)
 - Es soll möglich sein, Lesezeichen in Kategorien (z.B. Urlaub, Studium, JavaScript, Katzen, ...) zu organisieren.
 - Es soll möglich sein, neue Kategorien zu erzeugen.
 - Es soll möglich sein, die Zuordnung von Lesezeichen zu Kategorien zu ändern.
 - Lesezeichen sollen auch mehreren Kategorien zugeordnet werden können.
 - Kategorien sollen nicht hierarchisch strukturiert werden. 
 - Es soll möglich sein, nach Kategorien und in bestimmten Kategorien zu suchen.
 - Es soll möglich sein, Kategorien als sortierte Liste darzustellen.
 - Es soll möglich sein, nach Teilstrings aus Titel, URL oder Kurzbeschreibung zu suchen.
 - Automatisch erzeugte WVR-Kategorien (Watson Visual Recognition, siehe 2.4) und selbst erstellte Kategorien sollen separat behandelt werden können, z.B. Anzeige aller eigenen Kategorien oder nur der WVR-Kategorien, Suche in eigenen Kategorien oder nur in WVRKategorien usw. 
 - Aus Chrome exportierte Lesezeichen (HTML-Datei) sollen importiert werden können. Dabei soll mindestens die URL, der Linktext und die Kategorie übernommen werden. Die Übernahme der Icons ist optional. 
 - Eine Umkreissuche bezüglich der den Lesezeichen zugeordneten GPS-Positionen soll per Serveranfrage möglich sein. 
 - Die Umkreissuche soll wahlweise bezüglich eines eingegebenen Orts bzw. Position oder zu der Position eines anderen Lesezeichens ausgeführt werden können. 
 - Der Rich-Client soll per Ajax-Technik im Intervall von fünf Sekunden beim drückMich-Server anfragen, ob für diese Benutzer*in neue Daten vorliegen und dann evtl. die Seite aktualisieren. 

## drückMich-Server und Datenbank

 - Der drückMich-Server soll die Seiten für den drückMich-Client generieren, die Schnittstelle zur Datenbank bilden, die Lesezeichenseiten analysieren und Anfragen an externe Webdienste stellen.
 - Der Server soll aus jeder Lesezeichenseite eine Kurzbeschreibung und einen Titeltext extrahieren und abspeichern. 
 - Der Server soll für jede Lesezeichenseite mindestens ein Icon oder die URL eines Icons abspeichern. Optional kann für diese Analyse ein externer Webdienst verwendet werden. 
 - Der Server soll die image-URLs jeder Lesezeichenseite extrahieren. 
 - Falls in einer dieser Bilddateien eine GPS-Position gefunden wird, sollen diese Koordinaten mit dem Lesezeichen gespeichert werden (maximal eine Position pro Lesezeichen). 
 - Die Umkreissuche (siehe 2.3) soll als Datenbankoperation ausgeführt werden.
 - Aus einem oder aus mehreren Bildern einer Seite (falls vorhanden) sollen per Bildanalyse mit Hilfe des IBM Watson Visual Recognition (WVR) Dienstes automatisch Kategorien („classes“) für dieses Lesezeichen generiert und gespeichert werden.
 - Es sollen nur Bilder mit einer sinnvollen Mindestgröße analysiert werden.

# Implementationsanforderungen 
## Datenbank

 - Die drückMich-Datenbank soll mit dem DBMS MongoDB erstellt werden.
 - In der DB sollen alle Daten (Zugangsdaten, Sitzungsinformationen, Lesezeichendaten) verwaltet werden.
 - Zur Speicherung der Lesezeichendaten soll auch das GridFs verwendet werden (z.B. für IconImages). 
 - Dateien der Anwendung selbst (Images, CSS-, JS-Dateien, Template-Dateien, ...) sollen im Filesystem des Betriebssystems, nicht jedoch im GridFs der DB abgelegt werden.

## Sprachen, Werkzeuge, Techniken

 - Clientseitig soll verwendet werden:
   - HTML5, CSS3
   - JavaScript (ES6)/DOM-Scripting, HTML5 APIs nach Bedarf 
   - Ajax 
   - Chrome Browser Extension (https://developer.chrome.com/extensions)(siehe HinweiseZurHausarbeitWS18.pdf)
 - Die Client-Anwendung sollen für die aktuelle Version von Google Chrome ausgelegt sein. 
 - Die Client-Anwendung kann selbst erstellte SVG- oder Canvas-Grafiken, serverseitig mit Go generierte Grafiken oder HTML5 Custom Elements enthalten. 
 - Serverseitig soll verwendet werden:
   - die Sprache Go, Go-templates und die bei https://golang.org/ verfügbaren Pakete 
   - das Paket https://github.com/rwcarlsen/goexif zum Extrahieren von GPS-Daten aus Images (siehe HinweiseZurHausarbeitWS18.pdf)
   -  das IBM Watson Visual Recognition API (https://www.ibm.com/watson/services/visual-recognition/) zur Generierung von Kategorien aus Bildern (siehe auch HinweiseZurHausarbeitWS18.pdf) 
   - der Google S2 Converter (https://www.labnol.org/internet/get-favicon-imageof-websites-with-google/4404/) zur Icon-Suche oder alternativ ein anderer kostenloser Dienst (s. z.B. https://stackoverflow.com/questions/38599939/howto-get-larger-favicon-from-googles-api) 
   -  MongoDB mit dem mgo-Go-API (https://godoc.org/gopkg.in/mgo.v2) ODER globalsign/mgo (https://godoc.org/github.com/globalsign/mgo)
 - Es sollen keine zusätzlichen Bibliotheken oder Frameworks verwendet werden
 - Das Layout und die Gestaltung der Client-UIs wird nicht vorgegeben. Vielmehr soll die intuitive Bedienbarkeit im Vordergrund stehen.

# Anforderungen zum Abgabetermin 

 - Zur Präsentation am Abgabetermin soll:
   - das System auf einem eigenen Rechner lauffähig installiert sein 
   - die Datenbank auf dem eigenen Rechner oder alternativ auf dem HS-Server borsti installiert sein 
   - der Zugriff auf den oben genannten WVR-Dienst mit einem eigenen API-Key erfolgen 
   - in jedem Fall ein eindeutiger DB-Bezeichner folgender Form verwendet werden: HA18DB_vorname_nachname_matrnr, z.B.: HA18DB_donald_duck_42 
 - Zum Abgabetermin soll die gesamte Anwendung als ZIP-Datei hochgeladen werden (siehe Teilnahmebedingungen):
   -  Die ZIP-Datei vorname_nachname_matrnr_drueckMich.zip soll genau vier Ordner mit den Bezeichnern drueckMich, browserExtension, dump und doku enthalten. 
      - Die Anwendung soll sich im Ordner drueckMich befinden (alle Go-Quelldateien, Template-Dateien/Ordner, statischer Inhalt, JS, CSS usw.). 
      - Der Ordner browserExtension enthält die Chrome Browser Extension als lokal installierbare entpackte Erweiterung.
      - Ein mit mongodump erstellter MongoDB-Datenbestand soll sich in dump befinden. Dieser Datenbestand soll hinreichend sinnvoll sein, um alle Funktionen des Systems zu demonstrieren. 
      - Im Ordner doku sollen sich ausschließlich PDF-Dokumente befinden.
 - Wird die ZIP-Datei entpackt und der Ordner drueckMich in einen golang/src Ordner kopiert, so soll die Anwendung auf einem Windows- oder Unix-System compiliert und ausgeführt werden können. 
 - Der Mongo-Datenbestand soll mit mongorestore importiert werden können (siehe HinweiseZurHausarbeitWS18.pdf).
 - Die Anwendung im Ordner drueckMich soll eine Datenbankverbindung zu localhost:27017 (mongoDB default-Port) verwenden.
 - Der WVR-API-Key soll aus einer Textdatei apiKey.txt eingelesen werden, so dass ein anderer Key ohne Neucompilierung verarbeitet werden kann.
 - Die abgegebene Anwendung muss keinen gültigen WVR-API-Key enthalten, jedoch soll der Ort der apiKey.txt Datei und das erforderliche Eingabeformat des Keys auf dem Deckblatt der Dokumentation angegeben werden. 
 - Die borsti-DB soll eine sinnvolle Konfiguration für mindestens zwei Benutzer*innen enthalten. 
 - Die Zugangsdaten für das abgelieferte System sollen sein:
   - BenutzernameA: drueck Kennwort: mich 
   - BenutzernameB: push Kennwort: me

# Dokumentation 
 - Die Dokumentation im Ordner doku soll insbesondere beschreiben:
    - das UI-Konzept, das Layout, evtl. verwendete Grafik, Ajax-Kommunikation
    - die Struktur der Go-Anwendung (Dateien, Pakete, Funktionen, handler, templates, globale Daten und Datenstrukturen, ...) 
    - die Datenbankstruktur (mit Skizze)
    - die Analyse der externen Seiten (Kurzbeschreibung, Titeltext, URL, Icon, GPS-Position, Kategorien) 
    - den Import von Lesezeichen 
    - die Umkreissuche 
    - die Benutzer- und Zustandsverwaltung 
    - Probleme, Besonderheiten, Bemerkenswertes
    -  evtl. nicht erfüllte Anforderungen
  - Die Dokumentation soll aus Skizzen, Diagrammen, Bildern etc. und kurzen, verständlichen, gehaltvollen, textuellen Beschreibungen bestehen. Langatmiges Geschwafel führt zur Abwertung. 
  - Der Umfang der Text-Dokumentation (ohne Quellcode, Skizzen, Diagramme, Bilder) soll 5 A4 Seiten nicht übersteigen (Richtwert 400 Wörter/Seite).
  - Seiten ohne Header-Informationen (siehe Teilnahmebedingungen), loses Blattwerk, rückwärts auf dem Kopf eingeheftete und zusammen getackerte Seiten werden bei der Bewertung ignoriert. 
  - Auf dem Deckblatt der Textdokumentation soll unbedingt angegeben werden:
    - das verwendende Ziel-Betriebssystem (*) und der Browsertyp
    - die Startprozedur (Namen und Orte der zu startenden Go-Prozesse, URLs der UIs) 
    - Ort (Pfad, Dateiname, Zeilennummer) der Datenbankverbindungsdefinition 
    - Ort (Pfad, Dateiname) und ggfs. Format des WVR-API-Keys

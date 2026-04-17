# ✈️ Itinerary Prettifier CLI

> A command-line tool that converts raw flight itineraries into clean, customer-friendly formats.

---

## 📌 Overview

This project is a CLI tool built for "Anywhere Holidays" to bridge the gap between backend-generated flight itineraries and customer-ready outputs.

Administrators currently receive raw, text-based itineraries that are difficult for customers to read. This tool automates the transformation process by:

* Converting airport codes into human-readable names
* Formatting dates and times into friendly formats
* Cleaning up whitespace and formatting issues

---

## 🚀 Features

* ✨ Converts IATA (`#LAX`) and ICAO (`##EGLL`) codes into airport names
* 🏙️ (Bonus) Converts codes into **city names** using `*#CODE`
* 🕒 Formats ISO 8601 dates into readable formats:

  * `D(...)` → `DD Mmm YYYY`
  * `T12(...)` → 12-hour format with AM/PM
  * `T24(...)` → 24-hour format
* 🌍 Handles timezone offsets (including Zulu time `Z → +00:00`)
* 🧹 Normalizes vertical whitespace
* ❌ Robust error handling (invalid input, malformed CSV, etc.)
* 🎨 (Bonus) Optional colored CLI output

---

## 🛠️ Tech Stack

* Language: Go
* Standard Library:

  * `os` (file handling)
  * `bufio` / `io` (reading files)
  * `encoding/csv` (airport lookup parsing)
  * `time` (date formatting)
  * `strings` / `regexp` (text processing)

---

## 📂 Project Structure

```id="proj01"
itinerary-prettifier/
│
├── main.go              # Entry point
├── colors.go            # Color codes
├── io.go                # Input and output management
├── logic.go             # Core logic behind parsing process
├── validation.go        # Validation of arguments and required files
├── input.txt            # Example input
├── output.txt           # Example output
├── airport-lookup.csv   # Airport dataset
└── README.md
```

---

## ⚙️ Installation

```bash id="inst01"
# Clone the repository
git clone https://gitea.kood.tech/georgiisenotrusov/prettifier

# Navigate into the project
cd itinerary-prettifier
```

> Requires Go installed

---

## ▶️ Usage

```bash id="usage01"
go run . ./input.txt ./output.txt ./airport-lookup.csv
go run . ./input.txt ./output.txt ./airport-lookup.csv --stats
```

### Help flag

```bash id="usage02"
go run . -h
```

Output:

```
itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

---

## 🧠 How It Works

### Airport Codes

* `#XXX` → IATA code → Airport name
* `##XXXX` → ICAO code → Airport name
* `*#XXX` (Bonus) → City name

Airport data is dynamically loaded from a CSV file:

```
name, iso_country, municipality, icao_code, iata_code, coordinates
```

❗ Rules:

* Missing columns → **malformed**
* Empty values → **malformed**
* Unknown codes → left unchanged

---

### 📅 Date & Time Formatting

| Input Format                | Output Example     |
| --------------------------- | ------------------ |
| `D(2007-04-05T12:30−02:00)` | `05 Apr 2007`      |
| `T12(...)`                  | `12:30PM (-02:00)` |
| `T24(...)`                  | `12:30 (-02:00)`   |

Special case:

* `Z` → `(+00:00)`

Malformed inputs are left unchanged.

---

### 🧹 Whitespace Rules

* Convert `\v`, `\f`, `\r` → `\n`
* Max one empty line between content

---

## ❌ Error Handling

| Scenario            | Output                     |
| ------------------- | -------------------------- |
| Wrong arguments     | Show usage                 |
| Input file missing  | `Input not found`          |
| Lookup file missing | `Airport lookup not found` |
| Malformed CSV       | `Airport lookup malformed` |

> Output file is **not created or modified** on error.

---

## 🎁 Bonus Features

* 🌆 City name support (`*#CODE`)
* 🎨 Colored terminal output (highlight:

  * Airports
  * Cities
  * Dates & time offsets

---

## 🧪 Example

### Input

```
Flight from #LAX to ##EGLL
Departure: D(2024-03-10T14:30Z)
Time: T12(2024-03-10T14:30Z)
```

### Output

```
Flight from Los Angeles International Airport to London Heathrow Airport
Departure: 10 Mar 2024
Time: 02:30PM (+00:00)
```
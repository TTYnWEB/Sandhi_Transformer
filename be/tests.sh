#!/usr/bin/env bash

URL="http://localhost:8080/transform"

declare -A tests=(
  # # Consonant Sandhi
  # ["tat cit"]="ta[cc]it"
  # ["tat jana"]="ta[jj]ana"
  # ["sad dharma"]="sa[ddh]arma"
  # ["rājan candra"]="rāja[ñc]andra"
  # ["rājan jñāna"]="rāja[ñj]ñāna"
  # ["rājan ṭīkā"]="rāja[ṇṭ]īkā"
  # ["rājan tapaḥ"]="rāja[nt]apaḥ"
  # ["rājan patyate"]="rāja[mp]atyate"
  # ["tat śivaḥ"]="ta[cch]ivaḥ"
  # ["rāmas ca"]="rāma[śca]"
  # ["tat satyam"]="ta[s] satyam"
  # ["rāmas karma"]="rāma[sk]arma"
  # ["manas indra"]="mana[r]indra"
  #
  # # Vowel Sandhi
  # ["rama avatāra"]="ram[ā]vatāra"
  # ["kavi indra"]="kav[ī]ndra"
  # ["guru upadeśa"]="gur[ū]padeśa"
  # ["rāma īśvara"]="rām[ē]śvara"
  # ["rāma uttama"]="rām[ō]ttama"
  # ["rāma ēkaḥ"]="rām[ai]kaḥ"
  # ["rāma ōṣadhiḥ"]="rāma[u]ṣadhiḥ"
  #
  # # Visarga Sandhi
  # ["guruḥ patati"]="guru[f] patati"
  # ["rāmaḥ karoti"]="rāma[H] karoti"
  ["rāmaḥ āgacchati"]="rām[ō] gacchati" # NOTE: PROBLEM
  ["rāmaḥ śivaḥ"]="rāma[] śivaḥ"        # NOTE: PROBLEM
  ["rāmaḥ"]="rāma[ḥa]"                  # Depending on how you handle final visarga # NOTE: PROBLEM

  # Anusvara Sandhi
  # ["saṁ kara"]="sa[ṅ] kara"
  # ["saṁ gīta"]="sa[ṅ] gīta"
  # ["saṁ cit"]="sa[ñ] cit"
  # ["saṁ jñāna"]="sa[ñ] jñāna"
  # ["saṁ ṭīkā"]="sa[ṇ] ṭīkā"
  # ["saṁ tapaḥ"]="sa[n] tapaḥ"
  # ["saṁ darśana"]="sa[n] darśana"
  # ["saṁ pat"]="sa[m] pat"
  # ["saṁ bhāṣā"]="sa[m] bhāṣā"
)

for input in "${!tests[@]}"; do
  expected="${tests[$input]}"
  echo "🔹 Input: '$input'"
  response=$(curl -s -X POST -H "Content-Type: application/json" \
    -d "{\"text\": \"$input\"}" "$URL")

  echo "🔸 Expected to contain: '$expected'"
  echo "🔍 Response:"
  echo "$response" | jq .
  echo "----------------------------------------"
  echo "Press any key to continue..."
  read -n1 -s # -n1: read one character, -s: silent (do not echo)
done

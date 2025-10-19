#!/usr/bin/env bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Counters
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Function to run a single test
run_test() {
  local input="$1"
  local expected="$2"
  local description="$3"

  TOTAL_TESTS=$((TOTAL_TESTS + 1))

  echo -e "${BLUE}Test $TOTAL_TESTS:${NC} $description"
  echo -e "  Input: '$input'"
  echo -e "  Expected: '$expected'"

  # Run the transformer and capture output
  local actual=$(./main.mjs "$input" 2>/dev/null)
  local exit_code=$?

  # Check if the script executed successfully
  if [ $exit_code -ne 0 ]; then
    echo -e "  ${RED}ERROR: Script failed to execute${NC}"
    FAILED_TESTS=$((FAILED_TESTS + 1))
    echo
    return 1
  fi

  echo -e "  Actual: '$actual'"

  # Compare results
  if [ "$actual" = "$expected" ]; then
    echo -e "  ${GREEN}✓ PASS${NC}"
    PASSED_TESTS=$((PASSED_TESTS + 1))
  else
    echo -e "  ${RED}✗ FAIL${NC}"
    FAILED_TESTS=$((FAILED_TESTS + 1))
  fi
  echo
}

# Function to print summary
print_summary() {
  echo "================================================="
  echo -e "${BLUE}TEST SUMMARY${NC}"
  echo "================================================="
  echo -e "Total tests: $TOTAL_TESTS"
  echo -e "${GREEN}Passed: $PASSED_TESTS${NC}"
  echo -e "${RED}Failed: $FAILED_TESTS${NC}"

  if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}All tests passed! 🎉${NC}"
    exit 0
  else
    echo -e "${RED}Some tests failed. 😞${NC}"
    exit 1
  fi
}

# Check if main.mjs exists and is executable
if [ ! -f "./main.mjs" ]; then
  echo -e "${RED}Error: ./main.mjs not found${NC}"
  exit 1
fi

if [ ! -x "./main.mjs" ]; then
  echo -e "${YELLOW}Warning: ./main.mjs is not executable, attempting to run with node${NC}"
fi

echo -e "${BLUE}Starting Sanskrit IAST Sandhi Transformer Tests${NC}"
echo "================================================="
echo

# ===================================================================
# COMPREHENSIVE VISARGA SANDHI TESTS
# ===================================================================

# aḥ/āḥ + VOWELS (becomes o/ā + vowel)
run_test "rāmaḥ asti" "rāmo 'sti" "aḥ + a → o + avagraha"
run_test "devaḥ icchati" "devo icchati" "aḥ + i → o + i"
run_test "guruḥ upadeśa" "guro upadeśa" "aḥ + u → o + u"
run_test "rāmaḥ ṛtam" "rāmo ṛtam" "aḥ + ṛ → o + ṛ"
run_test "devaḥ eva" "devo eva" "aḥ + e → o + e"
run_test "guruḥ ojas" "guro ojas" "aḥ + o → o + o"
run_test "devāḥ asti" "devā asti" "āḥ + a → ā + a"
run_test "devāḥ icchanti" "devā icchanti" "āḥ + i → ā + i"

# aḥ/āḥ + VOICELESS STOPS (stays ḥ)
run_test "rāmaḥ karoti" "rāmaḥ karoti" "aḥ + k → no change"
run_test "devaḥ khaḍga" "devaḥ khaḍga" "aḥ + kh → no change"
run_test "guruḥ paśyati" "guruḥ paśyati" "aḥ + p → no change"
run_test "rāmaḥ phalam" "rāmaḥ phalam" "aḥ + ph → no change"
run_test "devāḥ karoti" "devāḥ karoti" "āḥ + k → no change"
run_test "devāḥ paśyanti" "devāḥ paśyanti" "āḥ + p → no change"

# aḥ/āḥ + VOICED STOPS (becomes o/ā)
run_test "devaḥ gacchati" "devo gacchati" "aḥ + g → o"
run_test "rāmaḥ ghaṭa" "rāmo ghaṭa" "aḥ + gh → o"
run_test "guruḥ jalati" "guro jalati" "aḥ + j → o"
run_test "devaḥ jhaṭiti" "devo jhaṭiti" "aḥ + jh → o"
run_test "rāmaḥ dadāti" "rāmo dadāti" "aḥ + d → o"
run_test "guruḥ dhavati" "guro dhavati" "aḥ + dh → o"
run_test "devaḥ bhāṣate" "devo bhāṣate" "aḥ + b → o"
run_test "rāmaḥ bhavet" "rāmo bhavet" "aḥ + bh → o"
run_test "devāḥ gacchanti" "devā gacchanti" "āḥ + g → ā"
run_test "devāḥ dadāti" "devā dadāti" "āḥ + d → ā"

# aḥ/āḥ + SIBILANTS (becomes matching sibilant)
run_test "rāmaḥ chatra" "rāmaś chatra" "aḥ + c → aś"
run_test "devaḥ chinnā" "devaś chinnā" "aḥ + ch → aś"
run_test "guruḥ śṛṇoti" "guraś śṛṇoti" "aḥ + ś → aś"
run_test "rāmaḥ ṣaḍ" "rāmaṣ ṣaḍ" "aḥ + ṣ → aṣ"
run_test "devaḥ sarvam" "devas sarvam" "aḥ + s → as"
run_test "devāḥ chatra" "devāś chatra" "āḥ + c → āś"
run_test "devāḥ śobhante" "devāś śobhante" "āḥ + ś → āś"

# aḥ/āḥ + DENTALS (becomes s)
run_test "rāmaḥ tatra" "rāmas tatra" "aḥ + t → as"
run_test "devaḥ thālī" "devas thālī" "aḥ + th → as"
run_test "devāḥ tatra" "devās tatra" "āḥ + t → ās"
run_test "devāḥ sthitāḥ" "devās sthitāḥ" "āḥ + sth → ās"

# aḥ/āḥ + NASALS AND LIQUIDS (becomes o/ā)
run_test "rāmaḥ nara" "rāmo nara" "aḥ + n → o"
run_test "devaḥ mama" "devo mama" "aḥ + m → o"
run_test "guruḥ yāti" "guro yāti" "aḥ + y → o"
run_test "rāmaḥ vada" "rāmo vada" "aḥ + v → o"
run_test "devaḥ ramati" "devo ramati" "aḥ + r → o"
run_test "guruḥ labhate" "guro labhate" "aḥ + l → o"
run_test "devāḥ nara" "devā nara" "āḥ + n → ā"
run_test "devāḥ mama" "devā mama" "āḥ + m → ā"

# aḥ/āḥ + ASPIRATED H (becomes o/ā)
run_test "rāmaḥ hanta" "rāmo hanta" "aḥ + h → o"
run_test "devāḥ hanta" "devā hanta" "āḥ + h → ā"

# ===================================================================
# NON-A VISARGA TESTS (iḥ, īḥ, uḥ, ūḥ)
# ===================================================================

# iḥ/īḥ + VOWELS (becomes ir/īr)
run_test "agniḥ asti" "agnir asti" "iḥ + a → ir"
run_test "agniḥ icchā" "agnir icchā" "iḥ + i → ir"
run_test "muniḥ upāsate" "munir upāsate" "iḥ + u → ir"
run_test "agniḥ eva" "agnir eva" "iḥ + e → ir"
run_test "muniḥ ojas" "munir ojas" "iḥ + o → ir"

# iḥ/īḥ + VOICELESS CONSONANTS (becomes iṣ/īṣ)
run_test "agniḥ karoti" "agniṣ karoti" "iḥ + k → iṣ"
run_test "muniḥ paśyati" "muniṣ paśyati" "iḥ + p → iṣ"
run_test "agniḥ tatra" "agniṣ tatra" "iḥ + t → iṣ"
run_test "muniḥ chatra" "muniṣ chatra" "iḥ + c → iṣ"
run_test "agniḥ ṭīkā" "agniṣ ṭīkā" "iḥ + ṭ → iṣ"
run_test "muniḥ śobhate" "muniṣ śobhate" "iḥ + ś → iṣ"
run_test "agniḥ ṣaṭ" "agniṣ ṣaṭ" "iḥ + ṣ → iṣ"
run_test "muniḥ sarvam" "muniṣ sarvam" "iḥ + s → iṣ"

# iḥ/īḥ + VOICED CONSONANTS (becomes ir/īr)
run_test "agniḥ gacchati" "agnir gacchati" "iḥ + g → ir"
run_test "muniḥ dadāti" "munir dadāti" "iḥ + d → ir"
run_test "agniḥ bhāti" "agnir bhāti" "iḥ + b → ir"
run_test "muniḥ jalati" "munir jalati" "iḥ + j → ir"

# iḥ/īḥ + NASALS/LIQUIDS (becomes ir/īr)
run_test "agniḥ nara" "agnir nara" "iḥ + n → ir"
run_test "muniḥ mama" "munir mama" "iḥ + m → ir"
run_test "agniḥ yāti" "agnir yāti" "iḥ + y → ir"
run_test "muniḥ vada" "munir vada" "iḥ + v → ir"
run_test "agniḥ ramati" "agnir ramati" "iḥ + r → ir"
run_test "muniḥ hanta" "munir hanta" "iḥ + h → ir"

# uḥ/ūḥ + VOWELS (becomes ur/ūr)
run_test "vasuḥ asti" "vasur asti" "uḥ + a → ur"
run_test "vasuḥ icchā" "vasur icchā" "uḥ + i → ur"
run_test "śakruḥ eva" "śakrur eva" "uḥ + e → ur"

# uḥ/ūḥ + VOICELESS CONSONANTS (becomes uṣ/ūṣ)
run_test "vasuḥ karoti" "vasuṣ karoti" "uḥ + k → uṣ"
run_test "vasuḥ paśyati" "vasuṣ paśyati" "uḥ + p → uṣ"
run_test "śakruḥ tatra" "śakruṣ tatra" "uḥ + t → uṣ"
run_test "vasuḥ chatra" "vasuṣ chatra" "uḥ + c → uṣ"
run_test "vasuḥ śobhate" "vasuṣ śobhate" "uḥ + ś → uṣ"

# uḥ/ūḥ + VOICED CONSONANTS (becomes ur/ūr)
run_test "vasuḥ gacchati" "vasur gacchati" "uḥ + g → ur"
run_test "vasuḥ dadāti" "vasur dadāti" "uḥ + d → ur"
run_test "śakruḥ bhāti" "śakrur bhāti" "uḥ + b → ur"

# uḥ/ūḥ + NASALS/LIQUIDS (becomes ur/ūr)
run_test "vasuḥ nara" "vasur nara" "uḥ + n → ur"
run_test "vasuḥ mama" "vasur mama" "uḥ + m → ur"
run_test "śakruḥ yāti" "śakrur yāti" "uḥ + y → ur"
run_test "vasuḥ hanta" "vasur hanta" "uḥ + h → ur"

# ===================================================================
# EDGE CASES AND COMPLEX EXAMPLES
# ===================================================================

# Multiple visargas in sequence
# run_test "rāmaḥ devaḥ asti" "rāmo devo 'sti" "Multiple visarga transformations"
# run_test "agniḥ vasuḥ karoti" "agniṣ vasuṣ karoti" "Mixed visarga types"
#
# # Visarga at end of phrase (no change)
# run_test "rāmaḥ" "rāmaḥ" "Single word with visarga"
# run_test "agniḥ" "agniḥ" "Single word with iḥ"
# run_test "vasuḥ" "vasuḥ" "Single word with uḥ"
#
# # Standalone visarga (should not change)
# run_test "ḥ asti" "ḥ asti" "Standalone visarga"#!/usr/bin/env bash

# # Basic visarga + vowel tests
# run_test "rāmaḥ asti" "rāmo 'sti" "Visarga + a → o + avagraha"
# run_test "devaḥ icchati" "deva icchati" "Visarga + i → dropped"
# run_test "guruḥ upadeśa" "guru upadeśa" "Visarga + u → dropped"
# run_test "rāmaḥ paśyati" "rāmaḥ paśyati" "Visarga + consonant → no change"
#
# # Vowel + vowel sandhi
# run_test "rāma asti" "rāmo 'sti" "a + a → o + avagraha"
# run_test "devi īśa" "devīśa" "i + ī → ī"
# run_test "guru upadeśa" "gurupadeśa" "u + u → u"
# run_test "rāma indra" "rāmendra" "a + i → e"
# run_test "deva upadeśa" "devopadeśa" "a + u → o"
# run_test "rāma ēva" "rāmaiva" "a + ē → ai"
# run_test "deva ōjas" "devauja" "a + ō → au"
#
# # More complex vowel combinations
# run_test "hari asti" "hary asti" "i + a → y + a"
# run_test "guru ātmā" "gurvātmā" "u + ā → v + ā"
# run_test "te ē" "taī" "e + ē → ai"
# run_test "yo ō" "yau" "o + ō → au"
#
# # Consonant + consonant sandhi
# run_test "tat śiva" "tacchiva" "t + ś → cch"
# run_test "rājan candra" "rājañcandra" "n + c → ñc"
# run_test "sad dharma" "saddharma" "d + dh → ddh"
# run_test "tat satyam" "atsatyam" "t + s → ts"
#
# # Anusvāra tests
# run_test "saṁ gacchati" "saṅgacchati" "ṁ + g → ṅg"
# run_test "saṁ tat cit" "saṁtaccit" "ṁ + t → ṁt, t + c → tc"
#
# # Edge cases
# run_test "rāmaḥ" "rāmaḥ" "Single word with visarga"
# run_test "ḥ asti" "ḥ asti" "Standalone visarga (should not change)"
# run_test "a" "a" "Single vowel"
# run_test "rāma" "rāma" "Word ending in vowel"
# run_test "rāmaḥ xyz" "rāmaḥ xyz" "Visarga before non-Sanskrit consonant"
# run_test "abc ḥ" "abc ḥ" "Non-Sanskrit + visarga"
# run_test "aḥ aḥ" "aḥ aḥ" "Multiple visargas"
# run_test "ṁ" "ṁ" "Standalone anusvāra"
# run_test "ṁ k" "ṅk" "anusvāra + k → ṅk"
#
# # Compound words (should remain unchanged if already in sandhi)
# run_test "saṁskāra" "saṁskāra" "Already combined compound"
# run_test "saṁpūrṇa" "saṁpūrṇa" "Already combined compound"
#
# # Longer phrases
# run_test "rāmaḥ asti sundara puruṣaḥ" "rāmo 'sti sundara puruṣaḥ" "Multiple word phrase"
# run_test "mama ātmā ananta ānanda" "mamātmā anantānanda" "Chain of vowel sandhi"
# run_test "tat tvam asi iti upaniṣat" "tattvam asi ity upaniṣat" "Classical Sanskrit phrase"
#
# # Test invalid inputs (if your script handles them)
# run_test "" "" "Empty input"
# run_test "   " "   " "Whitespace only"
#
# # Optional: Test with special characters that might break parsing
# run_test "rāma|asti" "rāma|asti" "Non-standard separator"
#
# echo
# print_summary

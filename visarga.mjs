// rules/visarga.mjs

const isVowel = char => 'aāiīuūṛṝḷḹēō'.includes(char);

/**
 * Creates a standardized result object for sandhi transformations
 * 
 * @param {string} result - The transformed text after applying sandhi rules
 * @param {boolean} [combined=false] - Whether the transformation combines words into one
 * @returns {{result: string, combined: boolean}} Object containing transformation result and combination status
 */
const sandhiResult = (result, combined = false) => ({
  result,
  combined,
});

/**
 * Applies visarga sandhi rules between two words
 * 
 * @param {string} word1 - First word (must end with visarga ḥ)
 * @param {string} word2 - Second word  
 * @returns {Object|null} Sandhi result object or null if no rule applies
 */
function tryVisargaRules(word1, word2) {
  if (!word1 || !word2)
    return null;

  const lastIdxW1 = (word1.length - 1);
  const tailW1 = word1[lastIdxW1];
  const headW2 = word2[0];
  const baseW1 = word1.slice(0, -1);

  // Early return if last char is not visarga ḥ
  if (tailW1 !== 'ḥ')
    return null;

  // Rule 1: visarga before voiceless bilabial stops (p/ph) → 'f' (upadhmānīya)
  if (headW2 === 'p')
    return sandhiResult(baseW1 + 'f');

  // Rule 2: visarga before voiceless velar stops (k/kh) → 'H' (jihvamūlīya)  
  if (headW2 === 'k')
    return sandhiResult(baseW1 + 'H');

  // Rule 3: visarga before sibilants (ś/ṣ/s) → dropped
  if ((headW2 === 'ś') || (headW2 === 'ṣ') || (headW2 === 's'))
    return sandhiResult(baseW1);

  if (isVowel(headW2))
    return sandhiResult(baseW1);

  // NOTE: visarga before vowels drops, then vowel sandhi can be applied!
  // This case should be handled by vowel rules after visarga is dropped

  return null;
}

export default tryVisargaRules;

const isVowel = c => 'aāiīuūṛṝḷḹēō'.includes(c);
const isSonant = c => 'gjdbnmyvrh'.includes(c) // voiced consonants + nasals/liquids

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

  const i = (word1.length - 1); // word1 last char index
  const h = (word1.length - 2); // word1 index for getting penultimate char // last 2 chars
  const tailW1 = word1[i];
  const penultW1 = word1[h];
  const baseW1 = word1.slice(0, -1);
  const base2W1 = word1.slice(0, -2);
  const headW2 = word2[0];
  const restW2 = word2.slice(1);

  if (tailW1 !== 'ḥ')
    return null;

  // NOTE: Becomes o before a, otherwise a
  if (penultW1 === 'a') {
    if (headW2 === 'a')
      return sandhiResult(base2W1 + 'ō')
      // return sandhiResult(`${base2W1}ō ${restW2}`, true);
    else if (headW2 === 'c')
      return sandhiResult(baseW1 + 'ś')
    else if (headW2 === 't')
      return sandhiResult(baseW1 + 's')
    else if (isVowel(headW2))
      return sandhiResult(base2W1 + 'a');
    else if (isSonant(headW2))
      return sandhiResult(base2W1 + 'ō');
  }

  if (penultW1 === 'ā') {
    if (isVowel(headW2))
      return sandhiResult(baseW1);
    else if (isSonant(headW2))
      return sandhiResult(baseW1);
    else if (headW2 === 'c')
      return sandhiResult(baseW1 + 'ś')
    else if (headW2 === 't')
      return sandhiResult(baseW1 + 's')
  }

  // voiceless stops
  if (headW2 === 'k') {
    if ((penultW1 === 'i') || (penultW1 ==='u'))
      return sandhiResult(baseW1 + 'ṣ');
      // NOTE: iḥ/uḥ becomes iṣ/uṣ
      // NOTE: iḥ/uḥ - ??? - or is it becomes the sibilant matching the homorganic equivalent of the following hard consonant
      // NOTE: iḥ/uḥ - ??? - and is it becomes 'r' before soft consonant or vowels
    else
      return sandhiResult(baseW1 + 'H');
  }

  if (headW2 === 'p')
    return sandhiResult(baseW1 + 'f');

  // voiced stops
  if (isSonant(headW2))
    return sandhiResult(baseW1 + 'r');

  // sibilants
  if (headW2 === 'c')
    return sandhiResult(baseW1 + 'ś')

  // dentals
  if (headW2 === 't')
    return sandhiResult(baseW1 + 's')

  // vowels
  if (isVowel(headW2))
    return sandhiResult(baseW1 + 'r')

  return null;
}

export default tryVisargaRules;


/*
  // // aḥ + a/ā  -> o (combined?)
  // if ((tail2W1 === 'aḥ') && ((headW2 === 'a') || (headW2 === 'ā')))
  //   return sandhiResult(`${base2W1}ō ${restW2}`, true);
  //
  // // aḥ + g/j/d/b/n/m/y/v -> o
  // if ((tail2W1 === 'aḥ') && isSonant(headW2))
  //   return sandhiResult(`${base2W1}ō ${restW2}`, true);
  //
  // // NOTE: need to ELSE the rest to not match ah?
  //
  // // aḥ + any-other-vowel -> dropped
  // // could be covered by final drop?
  // // if ((tail2W1 === 'aḥ') && isVowel(headW2))
  // //   return sandhiResult(baseW1);
  // // āḥ + *any* vowel -> dropped
  // // could be covered by final drop?
  // // if ((tail2W1 === 'āḥ') && isVowel(headW2))
  // //   return sandhiResult(base2W1);
  //
  // // before unvoiced consonants (not including p/k)
  // // ḥ + t -> s
  // if (headW2 === 't')
  //   return sandhiResult(baseW1 + 's');
  //
  // // ḥ + c -> ś
  // if (headW2 === 'c')
  //   return sandhiResult(baseW1 + 'ś');
  //
  // // ḥ + ṭ -> ṣ
  // // if preceded by i/ī, u/ū occasionally becomes ṣ, and the following t then becomes ṭ
  // if (headW2 === 'ṭ')
  //   return sandhiResult(baseW1 + 'ṣ');
  //
  // // visarga before voiceless bilabial stops (p/ph) → 'f' (upadhmānīya)
  // if (headW2 === 'p')
  //   return sandhiResult(baseW1 + 'f');
  //
  // // visarga before voiceless velar stops (k/kh) → 'H' (jihvamūlīya)  
  // if (headW2 === 'k')
  //   return sandhiResult(baseW1 + 'H');
  //
  // // visarga before sibilants (ś/ṣ/s) → dropped
  // // if ((headW2 === 'ś') || (headW2 === 'ṣ') || (headW2 === 's'))
  // //   return sandhiResult(baseW1);
  //
  // // ḥ (except after aḥ/āḥ) + g/j/d/b/n/m/y/v/(vowels) -> r
  // if (isSonant(headW2) || isVowel(headW2))
  //   return sandhiResult(baseW1 + 'r');
  //
  // // aḥ + *any other* vowel -> dropped
  // // ḥ (preceded by *any other vowel) -> dropped
  // if (isVowel(headW2))
  //   return sandhiResult(baseW1);
  //
  // // NOTE: visarga before vowels drops, then vowel sandhi can be applied!
  // // This case should be handled by vowel rules after visarga is dropped
*/

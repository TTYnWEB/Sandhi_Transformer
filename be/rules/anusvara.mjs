// [DATA] //////////////////////////////////////////////////////////////////////
const nasalMap = {
  'k': 'ṅ',
  'g': 'ṅ',
  'ṅ': 'ṅ',
  'c': 'ñ',
  'j': 'ñ',
  'ñ': 'ñ',
  'ṭ': 'ṇ',
  'ḍ': 'ṇ',
  'ṇ': 'ṇ',
  't': 'n',
  'd': 'n',
  'n': 'n',
  'p': 'm',
  'b': 'm',
  'm': 'm',
  'y': 'ñ',
  'r': 'ṇ',
  'l': 'n',
  'v': 'm',
};


// [FUNC] //////////////////////////////////////////////////////////////////////
/**
 * Creates a standardized result object for sandhi transformations
 * 
 * @param {string} result - The transformed text after applying sandhi rules
 * @param {boolean} [combined=false] - Whether the transformation combines words into one
 * @returns {{result: string, combined: boolean}|null} Object containing transformation result and combination status
 */
const sandhiResult = (result, combined = false) => ({
  result,
  combined,
});

/**
 * Replaces the intra-word anusvāra (ṁ) in a given word with the appropriate (homorganic)
 * nasal consonant based on the character that follows it, using a predefined nasal mapping.
 *
 * @param {string} word - The input word possibly containing an anusvāra (ṁ).
 * @returns {string} The transformed word with anusvāra replaced, if applicable.
 */
const applyIntraWordAnusvara = word => {
  if (!(word.includes('ṁ')))
    return word;
  const i = word.indexOf('ṁ');
  const j = (i + 1);
  const nextChar = word[j];
  if (!nextChar)
    return word;
  const replacementChar = nasalMap[nextChar];
  if (!replacementChar)
    return word;
  const split = word.split('ṁ');
  const transformedWord = split.join(replacementChar);
  return transformedWord;
};

/**
 * Applies anusvāra transformation rules between two words by checking if the first word
 * ends in 'ṁ' and replacing it with an appropriate nasal consonant based on the
 * starting character of the second word.
 *
 * @param {string} word1 - The first word, potentially ending with an anusvāra ('ṁ').
 * @param {string} word2 - The second word, whose initial character determines the nasal replacement.
 * @returns {string|null} The transformed word1 after anusvāra substitution and further processing,
 *                        or null if no transformation is applicable.
 */
function tryAnusvaraRules(word1, word2) {
  if (!word1 || !word2)
    return null;
  const i = (word1?.length - 1);
  const tailCharWord1 = word1[i];
  if (tailCharWord1 !== 'ṁ')
    return null;
  const headCharWord2 = word2[0];
  const replacementChar = nasalMap[headCharWord2];
  if (!replacementChar)
    return null;
  const transformed = (word1.slice(0, i) + replacementChar);
  // console.log({ headCharWord2, replacementChar, transformed });
  return sandhiResult(transformed);
}

// [EXPs] //////////////////////////////////////////////////////////////////////
export {
  applyIntraWordAnusvara,
  tryAnusvaraRules,
};

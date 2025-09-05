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

const applyIntraWordAnusvara = word => {
  if (!word.includes('ṁ'))
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

function tryAnusvaraRules(word1, word2) {
  const tailIdx = (word1.length - 1);
  const tailCharWord1 = word1[tailIdx];
  if (tailCharWord1 !== 'ṁ')
    return null;
  const headCharWord2 = word2[0];
  const replacementChar = nasalMap[headCharWord2];
  const transformed = (word1.slice(0, tailIdx) + replacementChar);
  // console.log({ headCharWord2, replacementChar, transformed });
  return sandhiResult(transformed);
}

export {
  applyIntraWordAnusvara,
  tryAnusvaraRules,
};

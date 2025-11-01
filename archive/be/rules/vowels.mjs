// [ DATA ] ////////////////////////////////////////////////////////////////////
const savarnaDirga = {
  'a+a': 'ā',
  'a+ā': 'ā',
  'ā+a': 'ā',
  'ā+ā': 'ā',
  'i+i': 'ī',
  'ī+i': 'ī',
  'i+ī': 'ī',
  'ī+ī': 'ī',
  'u+u': 'ū',
  'u+ū': 'ū',
  'ū+u': 'ū',
  'ū+ū': 'ū',
  // 'ē+ē': 'ē',
  // 'ō+ō': 'ō',
  'ṛ+ṝ': 'ṝ',
  'ṝ+ṛ': 'ṝ',
  'ḷ+ḹ': 'ḹ',
  'ḹ+ḷ': 'ḹ',
};

const guna = {
  'a+i': 'ē',
  'a+ī': 'ē',
  'ā+i': 'ē',
  'ā+ī': 'ē',
  'a+u': 'ō',
  'a+ū': 'ō',
  'ā+u': 'ō',
  'ā+ū': 'ō',
};

const vriddhi = {
  'a+ē': 'ai',
  'ā+ē': 'ai',
  'a+ō': 'au',
  'ā+ō': 'au',
  'a+ṛ': 'ar',
  'a+ṝ': 'ar',
  'ā+ṛ': 'ar',
  'ā+ṝ': 'ar',
  'a+ḷ': 'al',
  'a+ḹ': 'al',
  'ā+ḷ': 'al',
  'ā+ḹ': 'al',
};

// purvarupa 
// a word ending in e or o is followed by a word beginning with a.
// In this case, the a is simply elided (dropped), and the e or o remains unchanged.
const purvaRupa = {
  'ē+a': true,
  'ē+ā': true,
  'ō+a': true,
  'ō+ā': true,
};

// semi-vowel?
const yan = {
  'i': 'y',
  'ī': 'y',
  'u': 'v',
  'ū': 'v',
  'ṛ': 'r',
  'ṝ': 'r',
  'ḷ': 'l',
  'ḹ': 'l',
};

// kind of like yan extended?
const ayava = {
  'ē': 'ay',
  'ō': 'av',
};

const ayava2 = {
  'ai': 'āy',
  'au': 'āv',
};


// [ FUNC ] ////////////////////////////////////////////////////////////////////
// NOTE: must be *dissimilar* vowel?
// this might be covered by savarna already...
// since there they are *similar*
const isVowel = c => 'aāiīuūṛṝḷḹēō'.includes(c);

const sandhiResult = (result, combined = true) => ({ // NOTE: combined = true because all vowel rules combine
  result,
  combined,
});


// [ MAIN ] ////////////////////////////////////////////////////////////////////
function tryVowelRules(word1, word2) {
  if (!word1 || !word2)
    return null;

  const i = (word1.length - 1); // word1 last char index
  const j = (word1.length - 2); // index for getting last 2 chars of word2 
  const tailW1 = word1[i];
  const tail2W1 = word1.slice(j);
  const headW2 = word2[0];
  const baseW1 = word1.slice(0, -1);
  const base2W1 = word1.slice(0, -2);
  const restW2 = word2.slice(1);

  const key = `${tailW1}+${headW2}`;

  // Check savarna dirga first
  if (savarnaDirga[key]) {
    const replacement = savarnaDirga[key];
    return sandhiResult(baseW1 + replacement + restW2);
  }

  // Check guna
  if (guna[key]) {
    const replacement = guna[key];
    return sandhiResult(baseW1 + replacement + restW2);
  }

  // Check vriddhi
  if (vriddhi[key]) {
    const replacement = vriddhi[key];
    return sandhiResult(baseW1 + replacement + restW2);
  }

  // purvarupa 
  if (purvaRupa[key])
    return sandhiResult(`${word1}'${restW2}`);

  // Check yan
  if (yan[tailW1] && isVowel(headW2)) {
    const replacement = (yan[tailW1] + headW2);
    return sandhiResult(baseW1 + replacement + restW2);
  }

  // Check ayava
  if (ayava[tailW1] && isVowel(headW2)) {
    const replacement = (ayava[tailW1] + headW2);
    return sandhiResult(baseW1 + replacement + restW2);
  }

  if (ayava2[tail2W1] && isVowel(headW2)) {
    const replacement = (ayava2[tail2W1] + headW2);
    return sandhiResult(base2W1 + replacement + restW2);
  }

  // NOTE: "śivaḥ te ēva" → "śiva tyaiva" (ḥ drops, ē+a→ya, a+ē→ai)

  return null;
}

// [ EXPs ] ////////////////////////////////////////////////////////////////////
export default tryVowelRules;

const isVowel = c => 'aāiīuūṛṝḷḹēō'.includes(c);

const sandhiResult = (result, combined = false) => ({
  result,
  combined,
});

// Helper maps for vowels and their sandhi results
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
  'ē+ē': 'ē',
  'ō+ō': 'ō',
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
};

// Yaṇ sandhi: i/ī/ē + vowel → y + vowel, u/ū/ō + vowel → v + vowel
const yan = {
  'i': 'y',
  'ī': 'y',
  'ē': 'y',
  'u': 'v',
  'ū': 'v',
  'ō': 'v',
};

function tryVowelRules(word1, word2) {
  if (!word1 || !word2)
    return null;

  const lastIdxW1 = (word1.length - 1);
  const tailW1 = word1[lastIdxW1];
  const headW2 = word2[0];
  const baseW1 = word1.slice(0, -1);
  const restW2 = word2.slice(1);

  const key = `${tailW1}+${headW2}`;

  // Check savarna dirga first
  if (savarnaDirga[key]) {
    const replacement = savarnaDirga[key];
    return sandhiResult((baseW1 + replacement + restW2), true);
  }

  // Check guna
  if (guna[key]) {
    const replacement = guna[key];
    return sandhiResult((baseW1 + replacement + restW2), true);
  }

  // Check vriddhi
  if (vriddhi[key]) {
    const replacement = vriddhi[key];
    return sandhiResult((baseW1 + replacement + restW2), true);
  }

  // Check yan
  if (yan[tailW1] && isVowel(headW2)) {
    const replacement = yan[tailW1] + headW2;
    return sandhiResult((baseW1 + replacement + restW2), true);
  }

  // NOTE: "śivaḥ te ēva" → "śiva tyaiva" (ḥ drops, ē+a→ya, a+ē→ai)

  return null;
}

export default tryVowelRules;

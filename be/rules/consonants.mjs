// [ FUNC ] ////////////////////////////////////////////////////////////////////
const isVoiced = char => 'gjḍdbṅñṇnmyrlvh'.includes(char);
const isVowel = char => 'aāiīuūṛṝḷḹēō'.includes(char);
const sandhiResult = (result, combined = false) => ({
  result,
  combined,
});


// [ MAIN ] ////////////////////////////////////////////////////////////////////
const tryConsonantRules = (word1, word2) => {
  if (!word1 || !word2)
    return null;

  const i = (word1.length - 1); // word1 last char index
  const tailW1 = word1[i];
  const headW2 = word2[0];
  const baseW1 = word1.slice(0, -1);
  const restW2 = word2.slice(1);

  // Priority 1: Multi-character sequences (most specific first)
  if (word2.length >= 2) {
    const headW2c2 = word2.slice(0, 2);
    const restW2c2 = word2.slice(2);
    // s + ca → śca
    if ((tailW1 === 's') && (headW2c2 === 'ca'))
      return sandhiResult(`${baseW1}śca${restW2c2}`, true);
    // d + dh → ddh (combined)
    if ((tailW1 === 'd') && (headW2c2 === 'dh'))
      return sandhiResult(`${baseW1}ddh${restW2c2}`, true);
  }

  // Priority 2: s before voiced-consonants or vowels → r
  if ((tailW1 === 's') && (isVoiced(headW2) || isVowel(headW2)))
    return sandhiResult(`${baseW1}r`);

  // Priority 3: Voiceless stops before vowels (no transformation in sandhi application)
  // This would be for voicing, but in application we typically don't voice
  
  // Priority 4: Specific consonant assimilations
  switch (tailW1) {
    case 's':
      if (headW2 === 'k')
        return sandhiResult((baseW1 + 'sk' + restW2), true);
      break;
    case 'd':
      if (headW2 === 'k')
        return sandhiResult((baseW1 + 'tk' + restW2), true);
      else if (headW2 === 't')
        return sandhiResult((baseW1 + 'tt' + restW2), true);
      break;
    case 'n':
      switch (headW2) {
        case 'c': return sandhiResult((baseW1 + 'ñc' + restW2), true); break;
        case 'j': return sandhiResult((baseW1 + 'ñj' + restW2), true); break;
        case 'ṭ': return sandhiResult((baseW1 + 'ṇṭ' + restW2), true); break;
        case 't': return sandhiResult((baseW1 + 'nt' + restW2), true); break;
        case 'p': return sandhiResult((baseW1 + 'mp' + restW2), true); break;
        break;
      }
      break;
    case 't':
      switch (headW2) {
        case 'c': return sandhiResult((baseW1 + 'cc'  + restW2), true); break; // `${baseW1}cc${restW2}`;
        case 'j': return sandhiResult((baseW1 + 'jj'  + restW2), true); break;
        case 'g': return sandhiResult((baseW1 + 'dg'  + restW2), true); break;
        case 'k': return sandhiResult((baseW1 + 'tk'  + restW2), true); break;
        case 'ś': return sandhiResult((baseW1 + 'cch' + restW2), true); break;
        case 's': return sandhiResult((baseW1 + 's'));                  break;
        break;
      }
      break;
  }

  return null;
};


// [ EXPs ] ////////////////////////////////////////////////////////////////////
export default tryConsonantRules;

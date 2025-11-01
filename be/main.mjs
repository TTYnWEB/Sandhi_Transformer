#!/usr/bin/env node


// https://lrc.la.utexas.edu/eieol/vedol/110

//[ IMPs ]////////////////////////////////////////////////////////////////////
import normalize from './normalize.mjs';
import tryConsonantRules from './rules/consonants.mjs';
import tryVowelRules from './rules/vowels.mjs';
import tryVisargaRules from './rules/visarga.mjs';
import {
  applyIntraWordAnusvara,
  tryAnusvaraRules,
} from './rules/anusvara.mjs';

//[ ARGs ]////////////////////////////////////////////////////////////////////
const input = process.argv[2];

//[ FUNC ]////////////////////////////////////////////////////////////////////
const applySandhi = (inputTxt) => {
  console.log('applySandhi: ', { inputTxt });
  if (!inputTxt)
    return null;
  const normalized = normalize(inputTxt);
  const words = normalized.trim().split(' ').map(word => applyIntraWordAnusvara(word));
  const result = [];

  for (let i = 0; i < words.length; i++) {
    const j = (i + 1);
    const word1 = words[i];
    const word2 = words[j];

    if (!word2) {
      result.push(word1);
      break;
    }

    const processAttempt = (rulesArr, word1, word2) => {
      const advance = 1; // NOTE: default increment (by 1)
      const combined = false;
      let currWord = word1;

      for (const ruleFunc of rulesArr) {
        const attempt = ruleFunc(currWord, word2);
        if (attempt === null) // NOTE: no match, continue to next rule
          continue;

        const { result, combined } = attempt;

        if (combined)
          return { finalResult: result, combined  }; // NOTE: if combined, need to return immediately, no other rules can match...
        else
          currWord = result;
      }
      return { finalResult: currWord,  combined };
    };

    const rulesArr = [tryAnusvaraRules, tryVisargaRules, tryVowelRules, tryConsonantRules];
    const { finalResult, combined } = processAttempt(rulesArr, word1, word2);

    if (combined)
      words[j] = finalResult;
    else
      result.push(finalResult);
   }

  const resultStr = result.join(' ');
  const output = { outputTxt: resultStr };
  console.log('applySandhi: ', { output });
  return output;
};


//[ EXPs ]////////////////////////////////////////////////////////////////////
export default applySandhi;


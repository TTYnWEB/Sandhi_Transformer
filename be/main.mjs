#!/usr/bin/env node


// https://lrc.la.utexas.edu/eieol/vedol/110

// IMPs
import normalize from './normalize.mjs';
import tryConsonantRules from './rules/consonants.mjs';
import tryVowelRules from './rules/vowels.mjs';
import tryVisargaRules from './rules/visarga.mjs';
import {
  applyIntraWordAnusvara,
  tryAnusvaraRules,
} from './rules/anusvara.mjs';

// ARGs
const input = process.argv[2];

// FUNC
const applySandhi = input => {
  if (!input)
    return null;
  const normalized = normalize(input);
  const words = normalized.trim().split(' ').map(word => applyIntraWordAnusvara(word));
  const result = [];

  for (let i = 0; (i < words.length); i++) {
    const j = (i + 1);
    const word1 = words[i];
    const word2 = words[j];

    if (!word2) {
      result.push(word1); // NOTE: check for final visarga?
      break;
    }

    /////////////////////////////////////////////

    // const processAttempt = (ruleFunc, current = { word1 }, word2) => {
    const processAttempt = (rulesArr, word1, word2) => {
      for (const ruleFunc of rulesArr) {
        const attempt = ruleFunc(word1, word2);
        if (!attempt?.result)
          continue;
        const { result, combined } = attempt;
        if (result) {
          if (combined)
            return { finalResult: result, toPush: true };
          else 
            word1 = result;
        }
      }
      return { finalResult: word1, toPush: false };
    };

    const rulesArr = [tryAnusvaraRules, tryVisargaRules, tryVowelRules, tryConsonantRules];

    let { finalResult, toPush } = processAttempt(rulesArr, word1, word2);

    if (toPush)
        result.push(finalResult);
    else
      words[j] = finalResult;
      // continue;

    // if (!attempt)
    //   return word1;

////////////////////////////////////////////////////////////////


    // let current = { word1 };
    //
    // let attempt = tryAnusvaraRules(current.word1, word2);
    // //
    // if (attempt?.result)
    //   current = { word1: attempt.result };
    //
    // attempt = tryVisargaRules(current.word1, word2);
    //
    // if (attempt?.result)
    //   current = { word1: attempt.result };
    //
    //  attempt = tryVowelRules(current.word1, word2);
    //
    //  if (attempt?.result)
    //    current = { word1: attempt.result, word2 };
    //  else 
    //    attempt = tryConsonantRules(current.word1, word2);
    //
    // if (!attempt)
    //   result.push(word1);
    // else {
    //   if (attempt.combined)
    //     words[j] = attempt.result; // Words were combined - update words array for next iteration
    //   else
    //     result.push(attempt.result); // i++;
    // }
  };

  return result.join(' ');
};

// MAIN
const output = applySandhi(input);
console.log(output);
// console.log(JSON.stringify({ input, outpt }, null, 2));

// EXPs
export default applySandhi;

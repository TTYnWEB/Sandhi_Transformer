#!/usr/bin/env node

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

    let current = { word1 };

    let attempt = tryAnusvaraRules(current.word1, word2);
    
    if (attempt?.result)
      current = { word1: attempt.result };

    attempt = tryVisargaRules(current.word1, word2);

    if (attempt?.result)
      current = { word1: attempt.result };
    
    attempt = tryVowelRules(current.word1, word2);

    if (attempt?.result)
      current = { word1: attempt.result, word2 };
    else 
      attempt = tryConsonantRules(current.word1, word2);

    if (attempt) {
      if (attempt.combined)
        words[j] = attempt.result; // Words were combined - update words array for next iteration
      else
        result.push(attempt.result); // i++;
    } else {
      result.push(word1);
    }
  };

  return result.join(' ');
};

// MAIN
const output = applySandhi(input);
console.log({ input, output });
// console.log(JSON.stringify({ input, outpt }, null, 2));

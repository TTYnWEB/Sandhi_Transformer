// Simple Jest tests for Sanskrit sandhi rules

import { tryAnusvaraRules, applyIntraWordAnusvara } from '../rules/anusvara.mjs';
import tryVisargaRules from '../rules/visarga.mjs';
import tryVowelRules from '../rules/vowels.mjs';
import tryConsonantRules from '../rules/consonants.mjs';
// import normalize from '../normalize.mjs';
import applySandhi from '../main.mjs';

// describe('Normalize Function', () => {
//   test('converts to lowercase', () => {
//     expect(normalize('RĀMA')).toBe('rāma');
//   });
//
//   test('replaces colon with visarga', () => {
//     expect(normalize('rāma:')).toBe('rāmaḥ');
//   });
//
//   test('removes extra spaces', () => {
//     expect(normalize('rāma   gacchati')).toBe('rāma gacchati');
//   });
// });

describe('Anusvāra Rules', () => {
  test('transforms intra-word ṁ + k to ṅ', () => {
    expect(applyIntraWordAnusvara('saṁkara')).toBe('saṅkara');
  });

  test('transforms ṁ + k to ṅ between words', () => {
    const result = tryAnusvaraRules('taṁ', 'karoti');
    expect(result.result).toBe('taṅ');
    expect(result.combined).toBe(false);
  });

  test('transforms ṁ + c to ñ', () => {
    const result = tryAnusvaraRules('saṁ', 'ca');
    expect(result.result).toBe('sañ');
  });

  test('returns null when word does not end with ṁ', () => {
    expect(tryAnusvaraRules('rāma', 'karoti')).toBeNull();
  });
});

describe('Visarga Rules', () => {
  test('transforms aḥ + vowel to ō', () => {
    const result = tryVisargaRules('rāmaḥ', 'asti');
    expect(result.result).toBe('rāmō');
  });

  test('transforms aḥ + c to aś', () => {
    const result = tryVisargaRules('devaḥ', 'ca');
    expect(result.result).toBe('devaś');
  });

  test('transforms aḥ + t to as', () => {
    const result = tryVisargaRules('putraḥ', 'tiṣṭhati');
    expect(result.result).toBe('putras');
  });

  test('returns null when word does not end with ḥ', () => {
    expect(tryVisargaRules('rāma', 'asti')).toBeNull();
  });
});

describe('Vowel Rules', () => {
  test('combines a + ā to ā (savarna)', () => {
    const result = tryVowelRules('mama', 'ālayaḥ');
    expect(result.result).toBe('mamālayaḥ');
    expect(result.combined).toBe(true);
  });

  test('combines a + i to ē (guṇa)', () => {
    const result = tryVowelRules('mama', 'icchā');
    expect(result.result).toBe('mamēcchā');
  });

  test('combines a + ē to ai (vṛddhi)', () => {
    const result = tryVowelRules('mama', 'ēkaḥ');
    expect(result.result).toBe('mamaikaḥ');
  });

  test('elides a after ē (pūrva-rūpa)', () => {
    const result = tryVowelRules('te', 'atra');
    expect(result.result).toBe(`tē'tra`);
  });

  test('returns null for consonant endings', () => {
    expect(tryVowelRules('rāmat', 'asti')).toBeNull();
  });
});

describe('Consonant Rules', () => {
  test('transforms s + ca to śca', () => {
    const result = tryConsonantRules('haṁs', 'ca');
    expect(result.result).toBe('haṁśca');
    expect(result.combined).toBe(true);
  });

  test('transforms s + voiced consonant to r', () => {
    const result = tryConsonantRules('haṁs', 'gacchati');
    expect(result.result).toBe('haṁr');
  });

  test('transforms t + c to cc', () => {
    const result = tryConsonantRules('bhavat', 'ca');
    expect(result.result).toBe('bhavacca');
  });

  test('transforms d + k to tk', () => {
    const result = tryConsonantRules('mad', 'karoti');
    expect(result.result).toBe('matkaroti');
  });

  test('returns null when no rules apply', () => {
    expect(tryConsonantRules('rāma', 'asti')).toBeNull();
  });
});

describe('Edge Cases', () => {
  test('all functions handle null inputs', () => {
    expect(tryAnusvaraRules(null, 'word')).toBeNull();
    expect(tryVisargaRules(null, 'word')).toBeNull();
    expect(tryVowelRules(null, 'word')).toBeNull();
    expect(tryConsonantRules(null, 'word')).toBeNull();
  });

  test('all functions handle empty strings', () => {
    expect(tryAnusvaraRules('', 'word')).toBeNull();
    expect(tryVisargaRules('', 'word')).toBeNull();
    expect(tryVowelRules('', 'word')).toBeNull();
    expect(tryConsonantRules('', 'word')).toBeNull();
  });
});


describe('Full Sandhi Integration', () => {
  test('handles multi-rule transformations', () => {
    // visarga + vowel sandhi: rāmaḥ asti → rāmō + asti → rāmō'sti
    expect(applySandhi('rāmaḥ asti')).toBe("rāmō'sti");
    
    // anusvara + vowel: taṁ eva → taṅ eva (or similar based on implementation)
    expect(applySandhi('taṁ eva')).toBe('taṁ ēva');
    
    // multiple vowel rules: mama icchā asti → mamecchāsti  
    expect(applySandhi('mama icchā asti')).toBe('mamēcchāsti');
  });

  test('handles longer phrases', () => {
    // Three word phrase with multiple transformations
    expect(applySandhi('rāmaḥ ca sītā')).toBe('rāmaś ca sītā');
    
    // Four word phrase
    expect(applySandhi('mama ālaye devaḥ asti')).toBe("mamālaye devō'sti");
    
    // Mixed transformations
    expect(applySandhi('taṁ karoti mama icchā')).toBe('taṅ karoti mamecchā');
  });
});

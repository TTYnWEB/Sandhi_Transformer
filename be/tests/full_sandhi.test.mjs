import applySandhi from '../main.mjs';

//- NOTE: Integration tests use applySandhi and accept non-normalized input

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
    expect(applySandhi('mama ālaye devaḥ asti')).toBe("mamālayē dēvō'sti");
    
    // Mixed transformations
    expect(applySandhi('taṁ karoti mama icchā')).toBe('taṅ karōti mamēcchā');
  });
});

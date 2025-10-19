import normalize from '../normalize.mjs';

describe('Normalize Function', () => {
  test('converts to lowercase', () => {
    expect(normalize('RĀMA')).toBe('rāma');
  });

  test('replaces colon with visarga', () => {
    expect(normalize('rāma:')).toBe('rāmaḥ');
  });

  test('removes extra spaces', () => {
    expect(normalize('rāma   gacchati')).toBe('rāma gacchati');
  });
});


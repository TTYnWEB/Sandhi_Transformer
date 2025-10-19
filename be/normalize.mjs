const toLowerIAST = input => input.toLowerCase();

const replaceChar = (input, oldChar, newChar) => input.replaceAll(oldChar, newChar);

const removeNonIASTChars = input => {
  const re = /[^a-zāīūōēṛṝḷḹṁḥṅñṭḍṇśṣ ]/g;
  return input.replace(re, "");
};

const multiSpaceRE = /\s+/g; // new RegExp('\\s+', 'g');
const collapseSpaces = input => input.replace(multiSpaceRE, " ");

const normalize = s => {
  if (!s)
    return null;
  s = s.normalize('NFC');
  s = toLowerIAST(s);
  s = replaceChar(s, "ṃ", "ṁ");
  s = replaceChar(s, "o", "ō");
  s = replaceChar(s, "e", "ē");
  s = replaceChar(s, ":", "ḥ");
  s = removeNonIASTChars(s);
  s = collapseSpaces(s);
  return s;
};

export default normalize;

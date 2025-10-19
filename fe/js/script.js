//[ ELEM ]/////////////////////////////////////////////////////////////////////
const html        = document.documentElement;
const input       = document.getElementById('input');
const transBtn    = document.getElementById('trans-btn');
const output      = document.getElementById('output');
const outputLabel = document.getElementById('output-label');

//[ DATA ]/////////////////////////////////////////////////////////////////////
let prevVal = input.value;

//[ FUNC ]/////////////////////////////////////////////////////////////////////
const selectText = event => event.target.select();

const setTheme = mode => {
  html.setAttribute('data-theme', mode);
  localStorage.setItem('theme', mode);
};

const setInitialTheme = () => {
  const storedTheme = localStorage.getItem('theme');
  if (storedTheme) {
    html.setAttribute('data-theme', storedTheme);
    themeToggle.checked = (storedTheme === 'light');
  } else {
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
    html.setAttribute('data-theme', prefersDark ? 'dark' : 'light');
    themeToggle.checked = !prefersDark;
    localStorage.setItem('theme', prefersDark ? 'dark' : 'light');
  }
};

const transBtnClk = () => {
  const content = input.value;  
  transBtn.setAttribute('hidden', true);
  output.removeAttribute('hidden');
  outputLabel.removeAttribute('hidden');
  output.value = content;
  output.select();
};

const unHide = () => {
  output.removeAttribute('hidden');
  transBtn.setAttribute('hidden', true);
};

const inputChange = (event) => {
  const currVal = event?.target?.value;
  if (currVal !== prevVal) {
    transBtn.removeAttribute('hidden');
    output.setAttribute('hidden', true);
    outputLabel.setAttribute('hidden', true);
  }
};

//[ EventListeners ]////////////////////////////////////////////////////////////
transBtn.addEventListener('click', transBtnClk);
transBtn.addEventListener('touchstart', transBtnClk);
input.addEventListener('input', inputChange);
output.addEventListener('click', selectText);
output.addEventListener('touchstart', selectText);

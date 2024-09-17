const uKnow = document.querySelector('.suku-diketahui').value
const a = document.querySelector('.nilai-suku-diketahui').value
const b = document.querySelector('.beda').value
const uN = document.querySelector('.suku-n').value


const result = document.querySelector('.result')
const RBtn = document.querySelector('.rButton')


const rumusSequences = (suku, nilaiSuku, beda, sukuN) => {
    if(suku == 1) {
        return  nilaiSuku + (sukuN - 1) * beda
    }
}



RBtn.addEventListener('click', () => {
    result.innerHTML = rumusSequences(uKnow, a, b, uN)
    console.log(rumusSequences(uKnow, a, b, uN))
    console.log(Number(uKnow))
    console.log(Number(a))
    console.log(Number(b))
    console.log(uN)
})
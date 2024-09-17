const uKnow = document.querySelector('.suku-diketahui')
const a = document.querySelector('.nilai-suku-diketahui')
const b = document.querySelector('.beda')
const uN = document.querySelector('.suku-n')


const result = document.querySelector('.result')
const RBtn = document.querySelector('.rButton')


const rumusSequences = (suku, nilaiSuku, beda, sukuN) => {
    if(suku == 1) {
        return  nilaiSuku + (sukuN - 1) * beda
    }
}



RBtn.addEventListener('click', () => {

    const val = rumusSequences(Number(uKnow.value), Number(a.value), Number(b.value), Number(uN.value))

    result.innerHTML = `Suku ke ${uN.value} adalah ${val}`
})
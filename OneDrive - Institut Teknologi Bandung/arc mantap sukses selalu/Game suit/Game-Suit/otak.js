const textPlayer = document.querySelector("#textPlayer");
const textComputer = document.querySelector("#textComputer");
const textHasil = document.querySelector("#textHasil");
const textSkor = document.querySelector("#textSkor")
const splash = document.querySelector('.splash')
const tombolReset = document.querySelector('#tombolReset')
const gambarPilihan = document.querySelectorAll('.GambarPilihan')
let player;
let computer;
let hasil;
var skorPlayer = 0;
var skorKomputer = 0;

tombolReset.addEventListener("click", () =>{
    skorPlayer = 0
    skorKomputer = 0
    textSkor.textContent = "0 : 0"
})

gambarPilihan.forEach(iconify => iconify.addEventListener("click", () => {
    player = iconify.id;
    computerTurn();
    textPlayer.textContent = player;
    textComputer.textContent = computer;
    textHasil.textContent = cekPemenang();
    textSkor.textContent = setelSkor();
}));

function computerTurn(){
    const randNum = Math.floor(Math.random() * 3) + 1;

    switch(randNum){
        case 1:
            computer = "Batu";
            break;
        case 2:
            computer = "Gunting";
            break;
        case 3:
            computer = "Kertas";
            break;
    }
}

function cekPemenang(){
    if(player == computer){
        return "Seri"
    } else if (computer == "Batu") {
        return (player == "Kertas") ? "Kamu Menang" : "Kamu Kalah";
    } else if (computer == "Kertas") {
        return (player == "Gunting") ? "Kamu Menang" : "Kamu Kalah";
    } else if (computer == "Gunting") {
        return (player == "Batu") ? "Kamu Menang" : "Kamu Kalah";
}}

function setelSkor(){
    if (textHasil.textContent == "Kamu Menang"){
        skorPlayer++;
    } else if (textHasil.textContent == "Kamu Kalah"){
        skorKomputer++;
    }
    return `${skorPlayer} : ${skorKomputer}`
}

function hilang(){
    splash.classList.add('display-none');
}
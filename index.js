async function main(){
    let array = []
    const modal1 = document.getElementById('hrg').value;
    const dp1 = document.getElementById('dp').value;
    const cicilan1 = document.getElementById('cicil').value;

    for(let a=1; a<=5; a++){
        const in1 = document.getElementById("npv-"+a)
        array.push(in1.value)
        console.log(array[a-1])
    }

    const modal_f = Number(modal1)
    const dp_f = Number(dp1)
    const cicilan_f = Number(cicilan1)

    let data = {
        modal:modal_f,
        dp:dp_f,
        cicilan:cicilan_f,
        persen:[
            {inflasi:Number(array[0])},
            {inflasi:Number(array[1])},
            {inflasi:Number(array[2])},
            {inflasi:Number(array[3])},
            {inflasi:Number(array[4])}
        ]
    }
    console.log(data)
    await fetch("http://localhost:8080/api/npv",{
        mode:'no-cors',
        method:"POST",
        headers:{
            'Content-Type':'application/json',
        },
        body: JSON.stringify(data)
    })
    .then((Response)=>Response.text())
    .then(result => {
        alert("Succes Post Data:)")
    })
    Get()
}

async function Get(){
   await fetch("http://localhost:8080/api/npv",
   {
    method:"GET",
    headers:{
        "Access-Control-Allow-Origin":"http://localhost:8080"
    }
   })
    .then((response)=>response.json())
    .then((result)=>{
        console.log(result)
        // const tbody = document.getElementById("tbody");
        for(let a=1; a<=5; a++){
            // let str = "input"+a
            
            let in1 = document.getElementById("np-"+a)
            in1.innerText = result.Result.npv[a-1]
        }
        const irr = document.getElementById("irr")
        irr.innerText=result.Result.irr+"%"
    })
    handleClick(1);
}

async function handleClick(value) {
    let end = 0;
    let start = 0;
    let no = 1
    let sumPv = 0
    const sum = document.getElementById("n-t")

    await fetch("http://localhost:8080/api/npv",
   {
    method:"GET",
    headers:{
        "Access-Control-Allow-Origin":"http://localhost:8080"
    }
   })
    .then((response)=>response.json())
    .then((result)=>{
        console.log(result)
        // const tbody = document.getElementById("tbody");
        for(let a=1; a<=5; a++){
            // let str = "input"+a
            
            let in1 = document.getElementById("np-"+a)
            in1.innerText = result.Result.npv[a-1]
        }
        const irr = document.getElementById("irr")
        irr.innerText=result.Result.irr+"%"

    if(value === 1){
        end = 11;
        start=1;
        document.getElementById('title').innerHTML = "Lembar Perhitungan 1";
        
        var inputText = document.getElementById('npv-1').value;
        document.getElementById('outputNpv-1').innerHTML = inputText;
        
        var koma = ((inputText/100)/12).toFixed(8);
        document.getElementById('outNpv-1').innerHTML = koma;

        for(let a=start; a<=end; a++){
            let str = "n-" + no
            let n = document.getElementById(str)
            console.log(result.Result.pv[a-1])
            n.innerText = result.Result.pv[a-1]
            sumPv += result.Result.pv[a-1]
            no++
        }
        sum.innerText = sumPv

    }else if(value === 2){
        end = 22;
        start=12;
        document.getElementById('title').innerHTML = "Lembar Perhitungan 2";
        
        var inputText = document.getElementById('npv-2').value;
        document.getElementById('outputNpv-1').innerHTML = inputText;
        
        var koma = ((inputText/100)/12).toFixed(8);
        document.getElementById('outNpv-1').innerHTML = koma;

        for(let a=start; a<=end; a++){
            let str = "n-" + no
            let n = document.getElementById(str)
            n.innerText = result.Result.pv[a-1]
            sumPv += result.Result.pv[a-1]
            no++
        }
        sum.innerText = sumPv
    }else if(value === 3){
        end = 33;
        start=23;
        document.getElementById('title').innerHTML = "Lembar Perhitungan 3";
        
        var inputText = document.getElementById('npv-3').value;
        document.getElementById('outputNpv-1').innerHTML = inputText;
        
        var koma = ((inputText/100)/12).toFixed(8);
        document.getElementById('outNpv-1').innerHTML = koma;

        for(let a=start; a<=end; a++){
            let str = "n-" + no
            let n = document.getElementById(str)
            console.log(result.Result.pv[a-1])
            n.innerText = result.Result.pv[a-1]
            sumPv += result.Result.pv[a-1]
            no++
        }
        sum.innerText = sumPv
    }else if(value === 4){
        end = 44;
        start = 34;
        document.getElementById('title').innerHTML = "Lembar Perhitungan 4";
        
        var inputText = document.getElementById('npv-4').value;
        document.getElementById('outputNpv-1').innerHTML = inputText;
        
        var koma = ((inputText/100)/12).toFixed(8);
        document.getElementById('outNpv-1').innerHTML = koma;

        for(let a=start; a<=end; a++){
            let str = "n-" + no
            let n = document.getElementById(str)
            console.log(result.Result.pv[a-1])
            n.innerText = result.Result.pv[a-1]
            sumPv += result.Result.pv[a-1]
            no++
        }
        sum.innerText = sumPv
    }else if(value === 5){
        end = 55;
        start=45;
        document.getElementById('title').innerHTML = "Lembar Perhitungan 5";
        
        var inputText = document.getElementById('npv-5').value;
        document.getElementById('outputNpv-1').innerHTML = inputText;
        
        var koma = ((inputText/100)/12).toFixed(8);
        document.getElementById('outNpv-1').innerHTML = koma;

        for(let a=start; a<=end; a++){
            let str = "n-" + no
            let n = document.getElementById(str)
            console.log(result.Result.pv[a-1])
            n.innerText = result.Result.pv[a-1]
            sumPv += result.Result.pv[a-1]
            no++
        }
        sum.innerText = sumPv
    }
    
    })

}
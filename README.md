# Dansk
dansk er et simpel og intuitivt programmeringssprog, der er designet til at gøre koding mere tilgængeligt ved at bruge dansk ordforråd og syntaks. Sproget understøtter grundlæggende programmeringselementer som variabler, funktioner, løkker og rækker. Det er perfekt til begyndere, men kan også bruges til mere komplekse opgaver.

# Eksempler på kode i *dansk*

### Eksempel 1: Summen af de første n tal

```
funktion sum_tal(n) {
    hvis (n <= 0) {udlever 0;}
    udlever n + sum_tal(n - 1);
}

meddel(sum_tal(5));
```

Dette eksempel beregner summen af de første 5 tal ved hjælp af en rekursiv funktion. Hvis `n` er 5, vil det udskrive summen af 1 + 2 + 3 + 4 + 5.

### Eksempel 2: Oprettelse og udskrift af en række

```
lad liste = [10. 20. 30. 40];

meddel(liste[1]);
```

Her defineres en række med fire tal. Programmet udskriver det andet element i rækken, som vil være `20`.

### Eksempel 3: Løkke der tæller ned

```
lad tæller = 10;
imens (tæller > 0) {
    meddel(tæller);
    tæller -= 1;
}
```

Dette eksempel viser en løkke, der starter fra 10 og tæller ned til 1. På hvert trin udskriver programmet værdien af `tæller`.

### Eksempel 4: Udskriv Beskeder

```
meddel("Hej Verden");
```



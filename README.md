# Dansk
dansk er et simpel og intuitivt programmeringssprog, der er designet til at gøre koding mere tilgængeligt ved at bruge dansk ordforråd og syntaks. Sproget understøtter grundlæggende programmeringselementer som variabler, funktioner, løkker og rækker. Det er perfekt til begyndere, men kan også bruges til mere komplekse opgaver.

# Grundlæggende syntaks i dansk
Her er en kort forklaring af de grundlæggende syntaksregler i dansk.

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

### Eksempel 2: Er et tal deleligt med 4 og 6?

```
funktion er_delelig(tal) {
    hvis (tal <= 0) {
        udlever falsk;
    }
    imens (tal <= 100) {
        hvis (tal % 4 == 0 og tal % 6 == 0) {
            udlever sandt;
        } 
        ellers {
            tal += 1;
        }
    }
    udlever falsk;
}

meddel(er_delelig(12));
```

Denne funktion tjekker, om et tal er deleligt med både 4 og 6. Hvis det er tilfældet, returneres `sand`; ellers `falsk`.

### Eksempel 3: Oprettelse og udskrift af en række

```
lad liste = [10. 20. 30. 40];

meddel(liste[1]);
```

Her defineres en række med fire tal. Programmet udskriver det andet element i rækken, som vil være `20`.

### Eksempel 4: Løkke der tæller ned

```
lad tæller = 10;
imens (tæller > 0) {
    meddel(tæller);
    tæller -= 1;
}
```

Dette eksempel viser en løkke, der starter fra 10 og tæller ned til 1. På hvert trin udskriver programmet værdien af `tæller`.

### Eksempel 5: Udskriv Beskeder

```
meddel("Hej Verden");
```



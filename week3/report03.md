### Improvement of Lab1,2
Since Lab1 already supports decryption, all I had to do was add a brute force option to it. If the limit `-l` is set to `-1`, then the program will output all possible decryptions. 

#### 1A
![](/assets/3-1.png)

#### 1B
![](/assets/3-3.png)

As for Lab2, I had to add a decryption option altogether. If the flag `-d` is passed and the flag `-k` is set to an integer < 26, then the input string `-t` is decrypted. If the flag `-k` is not set, then all 25 possible decryption shifts are given.

![](/assets/3-2.png)

### Lab3 
The main goal of Lab3 is to write a program that will brute force all possible decryptions of Lab1 and Lab2 and find a solution. So, I just combined the decryption algorithms from Lab1 and Lab2, where if you pass a string `-t`, it will be decrypted in all possible ways. 

![](/assets/3-4.png)

In the final version, I modified it so I don't print so much that the program is faster. Now, two arrays are created, populated by possible limits and shifts for decryptions.

![](/assets/3-6.png)

### Imrovements
Possible improvements could include the addition of multithreading. It would also be beneficial to avoid creating the hashmap each time during brute-forcing, optimizing memory usage and execution speed.

### Bonus
XOP, also known as XOR, can be easily used in encryption. Let's take the plaintext `"Slizik"` and convert every letter to ASCII integers, then to binary.

**Plaintext:**

- `S` = 83 → `1010011`
- `l` = 108 → `1101100`
- `i` = 105 → `1101001`
- `z` = 122 → `1111010`
- `i` = 105 → `1101001`
- `k` = 107 → `1101011`

**Key: `"key"`**

- `k` = 107 → `1101011`
- `e` = 101 → `1100101`
- `y` = 121 → `1111001`

Since the key is shorter than the plaintext, we repeat the key.

### XOR Table Example:
| **A** | **B** | **A XOR B** |
|-------|-------|-------------|
| 1     | 1     | 0           |
| 1     | 0     | 1           |
| 0     | 1     | 1           |
| 0     | 0     | 0           |

**XOR Encryption:**
We XOR each binary digit of the plaintext with the corresponding key digit.

| **Plaintext (A)** | **Key (B)** | **A XOR B** |
|-------------------|-------------|-------------|
| `1010011` (S)     | `1101011`   | `0111000`   |
| `1101100` (l)     | `1100101`   | `0001001`   |
| `1101001` (i)     | `1111001`   | `0010000`   |
| `1111010` (z)     | `1101011`   | `0010001`   |
| `1101001` (i)     | `1100101`   | `0001100`   |
| `1101011` (k)     | `1111001`   | `0010010`   |

**Encrypted text (binary):** `0111000 0001001 0010000 0010001 0001100 0010010`

**Encrypted text (hexadecimal):** `380910110c12`

### XOR Decryption
Since XOR has a unique property that if you apply XOR to the same value twice, you get the original result. If you know the key, you can easily retrieve the information.

![](/assets/3-5.png)

"""
Very simple - native large integer support a lifesaver
"""

def main():
    sum = 0
    with open("./one_hundred_numbers.txt") as f:
        first = int(f.readline())
        sum += first
        for line in f.readlines():
            sum += int(line)
        
    print("Sum: ", sum)
            


if __name__ == '__main__':
    main()
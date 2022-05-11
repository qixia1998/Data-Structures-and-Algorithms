
# example to understand variables
a = [2, 4, 6]
b = a
a.append(8)
print(b)


# example to understand variables scope

a = 15; b = 25
def my_function():
    global a
    a = 11; b = 21

my_function()
print(a) #11
print(b) #25
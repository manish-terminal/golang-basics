1️⃣ Overall Structure

-rw-r--r-- me total 10 characters hain:

- r w -  r - -  r - -


1st character → file type

Next 9 characters → permissions (3 sets of 3)

2️⃣ 1st Character: File Type
Character	Meaning
-	Regular file
d	Directory
l	Symbolic link

Tumhare case: - → Regular file

3️⃣ Next 9 Characters: Permissions

Ye 3 groups me divided hain:

rw-  r--  r--

Group	Meaning
1st (owner)	rw- → owner ke permissions
2nd (group)	r-- → group ke permissions
3rd (others)	r-- → others ke permissions
4️⃣ Individual Permissions
Character	Meaning
r	Read (padh sakta hai)
w	Write (likh sakta hai / modify kar sakta hai)
x	Execute (run kar sakta hai)
-	Permission nahi hai
5️⃣ Tumhare case: -rw-r--r--
Entity	Permission	Explanation
Owner	rw-	Owner file padh sakta hai aur modify kar sakta hai, execute nahi kar sakta
Group	r--	Group members file sirf padh sakte hain, modify ya execute nahi
Others	r--	Baaki sab file sirf padh sakte hain
6️⃣ Numeric Representation (Optional)

Linux me numeric form bhi hoti hai (chmod me use hoti hai):

rw- r-- r-- → 644


Owner: rw- = 4 (read) + 2 (write) + 0 (no exec) = 6

Group: r-- = 4

Others: r-- = 4

Command me:

chmod 644 filename

🔹 Summary

- → regular file

rw- → owner read/write

r-- → group read only

r-- → others read only

Numeric = 644

//USE bufio for byte reading writing straming fashion
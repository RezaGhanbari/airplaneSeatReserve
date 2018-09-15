### Simplest online airplane seat system

- Written in golang
- Dockerized 
- Postman collection added


###### environment variables :
- PORT

#### Description
1. The program should allow users to define the dimension of the
matrix that will represent seating arrangement.
Example: Input: Rows: 10, SeatsPerRow: 6

2. The program should allow users to select or input a seating assignment
randomly in the matrix or in sequence for a passenger:
Example: Input: Row: 3, Seat: 2, Occupied: true

3. The program should allow users to search for a seat number
by inputing the sequence of seat assigned.
Example:
Input: 3 (means 3rd seat assigned)
Output: 6F (depending on what part/coordinate in the matrix{seating arrangement}
the 3rd passenger was assigned.)

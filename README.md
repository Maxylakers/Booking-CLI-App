This application is a simple tickets booking application to be ran on CLI.

- It is written in Golang in a bid to test concurrency capabilities of Go.

- User data are collected and processed as a real world simulation.

- Collects user data like first-name, last-name, email, number-of-tickets to book.

- Further validates the collected data against provided conditions.

- If the data passes validation, it gets processed and stored in a slice. Also it returns the information below to user:

---Thank you *username* for booking *number-of-tickets* tickets. You will receive a confirmation email at       *user-email* 

We now have: *number-of-remaining-tickets* tickets remaining for the Max Conference

The first names of all bookings are:  [*username*]


- Proceeded to isolate Functions into packages so they can be ran using the:
go run . (to run all the files in the project, using the main.go at the entry point file)

- Simulate the tickets being sent to the users with a delay of 30 seconds.

- Ultimately used the go routine to make the application concurrent while sending the ticket is being simulated, thereby, not leaving our code blocked for the process to complete. 

- Proceeded to synchronize Go routines, such that when the main completes, and a green thread (a high level abstraction of a regular CPU thread) is still running, the application waits using waitGroups for the green threads to complete its task(s) before exiting.
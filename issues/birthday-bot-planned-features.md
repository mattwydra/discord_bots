## Birthday Bot - Feature: Register User Birthday

**Description**:  
Implement a command to register a user's birthday, verify the input, and store it in a database. The bot should send a confirmation message before proceeding.

**Tasks**:  
[ ] Add `/register` command to initiate the birthday registration process.  
[ ] Send DM to user asking them if they are sure they want to share their birthday info.  
[ ] Validate birthday format (month and day). If invalid, ask for re-entry.  
[ ] Ask user to confirm the details with a ✅ or ❌ reaction.  
[ ] Push confirmed birthday to the database with the associated user’s name or account.  
[ ] Optional: Ask for the user’s timezone for tracking.  
[ ] Daily check at midnight GMT to remind users of upcoming birthdays.

**Additional Notes**:  
Consider security implications and ensure that user data is stored privately. React appropriately to incorrect inputs and manage errors gracefully.  

---
## Birthday Bot - Feature: Themed Birthday Messages

**Description**:  
Add the ability to send themed or quirky birthday messages when a user’s birthday comes up. Users should be able to choose from a set of messages.

**Tasks**:  
- Provide users with a selection of themed birthday messages.  
- Store the selected message in the database alongside the birthday entry.  
- Send the selected birthday message on the day of the user’s birthday.

**Additional Notes**:  
Think about how to make the birthday messages fun and engaging, maybe add customization or randomization options.

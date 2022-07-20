# SlackBotGo
I have created Slack Bot, in this project i have used Go Language, and i have used the slacker library, 
Go to slack create your account create your workspace and give a name according to you.
After creating a slack account then go to api.slack.com for creating bot.
Then create a new app there from scratch/template
Then activate the socket mode, write token then generate socket token, and copy this token because i will paste that token in set environment.
Then go to the Oauth & permission Allow Bot Token Scopes 
    1 app_mention:read 
    2 channel:history
    3 channel:read
    4 chat:write
    5 im:read
    6 im:write
    7 mpim:hisotry
    8 mpim:read
    9 mpim:write
    
 Then we will find the event subscription option then add permission of events and save changes. subscribe the bot events
     1 app.mention
     2 message.im
     3 message.mpim
     4 message.group
     5 message.channel
     
  After this go to Oauth & permission  page then install the workspace, if you made any changes then re-install the workspace. 
 Write the code and after this run your code using " go run main.go" .

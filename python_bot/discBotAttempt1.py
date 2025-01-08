import discord
import certifi
import json
import requests
import urllib.parse
import random
from discord.ext import commands

client = commands.Bot(command_prefix = '?')
# can be '.' or '/' etc
# When bot is ready...
commandList = []
commandList.append('ping')
commandList.append('gn')
commandList.append('gm')
commandList.append('schedule')
commandList.append('calendar')
commandList.append('esports')
commandList.append('cpc')
commandList.append('rng')
commandList.append('gcd')
commandList.append('command_list')

@client.event
async def on_ready():
    print('Bot has connected to Discord!')

@client.event
async def on_member_join(member):
    print(f'{member} has joined the server :D')
    await message.channel.send(f'Hello, {member} :D')

@client.event
async def on_member_remove(member):
    print(f'so sad to see you go, {member}! Please come back soon! D:')
    await message.channel.send(f'Goodbye, {member} D:')

@client.command()
async def ping(ctx):
    await ctx.send(f'Pong! {round(client.latency * 1000)}ms')

@client.command(aliases = ['gn', 'Gn', 'GN', 'Goodnight'])
async def goodnight(ctx):
    await ctx.send('Goodnight!')

@client.command(aliases = ['gm', 'Gm', 'Goodmorning', 'GM'])
async def goodmorning(ctx):
    await ctx.send('Good morning! The current weather in Williamsburg for today is as follows:')

    base = 'http://api.openweathermap.org/data/2.5/weather?'
    app_id = '8b8c155b1466395e049c154916f169c1'
    city = 'Williamsburg'

    response = requests.get(base + 'q=' + city + '&appid=' + app_id)
    response = response.json()

    data = response["main"]
    currentTemperature = data["temp"]
    weather = response["weather"]
    description = weather[0]["description"]

    await ctx.send('temperature is: ' + str(round((currentTemperature - 273) * 9/5 + 32)) + ' degrees Fahrenheit, ' + 'or ' + str(round(currentTemperature - 273)) + ' degrees Celsius!')
    await ctx.send('current weather: ' + description)
    await ctx.send('Have a good day, everyone!')


@client.command()
async def schedule(ctx):
    str = 'https://docs.google.com/spreadsheets/d/1-FTbsAutbjqPNo8j3OI5A1nw8j1NU9w4NP5ElHqPKjs/edit#gid=991819808'
    await ctx.send('Hey, we dont need MEE6 the Discord Bot! Testing the spreadsheet link... beep beep boop boop... here you go! ' + str)

@client.command(aliases = ['cal', 'Cal', 'Calendar'])
@commands.has_role('Varsity Team')
async def calendar(ctx):
    cal1 = 'Practice Schedule: \nhttps://calendar.google.com/calendar/u/2?cid=Y19vbm82N2FvOTJoNjI3anFhZzZkcWRncnJlMEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t'
    cal2 = 'EGF Match Schedule: \nhttps://calendar.google.com/calendar/u/2?cid=Y19tMjJkbW9jc21vMGttcHRrOW1uNGgwZW10NEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t'
    cal3 = 'Team Meeting Schedule: \nhttps://calendar.google.com/calendar/u/2?cid=Y19qcTlyNWh1OGk1MTU1NWw2Mm82b2s1MGczZ0Bncm91cC5jYWxlbmRhci5nb29nbGUuY29t'
    await ctx.author.send('Here are the calendars:\n' + cal1 + cal2 + cal3)

@client.command()
async def esports(ctx):
    twitter = 'https://twitter.com/esportsatwm'
    instagram = 'https://www.instagram.com/esportsatwm/'
    stream1 = 'https://www.twitch.tv/officialegf'
    stream2 = 'https://www.twitch.tv/egfrocketleague'

    await ctx.send('Twitter: ' + twitter)
    await ctx.send('Instagram: ' + instagram)

    await ctx.send('Catch our games every Wednesday!')
    await ctx.send('Main stream: ' + stream1)
    await ctx.send('Occasionally here: ' + stream2)

@client.command(aliases = ['cpc', 'Copycat'])
async def copycat(ctx, *, strToCopy):
    await ctx.send(strToCopy)

@client.command(aliases = ['rand', 'Random', 'RNG', 'rng'])
async def randNumGen(ctx, minimum, maximum):
    min = int(minimum)
    max = int(maximum)
    num = random.choice(range(min, max + 1))
    await ctx.send(f'chose a number between {min} and {max} and got {num}!!')

@client.command(aliases = ['GCD', 'GreatestCommonDenominator', 'Euclid', 'Euclidean'])
async def gcd(ctx, num1, num2):
    n1 = int(num1)
    n2 = int(num2)
    gcd = 1
    if (n1 > n2):
        while (n2 != 0):
            temp = n2
            n2 = n1 % n2
            n1 = temp
        gcd = n1
    if (n2 > n1):
        while (n1 != 0):
            temp = n1
            n1 = n2 % n1
            n2 = temp
        gcd = n2
    if n1 == n2:
        gcd = n1
    await ctx.send(f'Greatest Common Denominator of {num1} and {num2} is: {gcd}')

@client.command(aliases = ['command_list', 'commands', 'Command_List', 'Commands'])
async def comList(ctx):
    await ctx.send(str(commandList))

@client.command(aliases = ['Fus', 'FUS', 'FUs', 'fUS', 'fUs', 'fuS'])
async def fus(ctx):
    await ctx.send("RO DAH!")
    await ctx.send('https://www.youtube.com/watch?v=AVy7YPNP_zI')



# @client.command(aliases = ['Prefix', 'pre', 'Pre'])
# async def prefix(ctx, newCommand):
#     client = commands.Bot(command_prefix = newCommand)

client.run('TOKEN')

import discord
import certifi
import json
import requests
import urllib.parse
import random
from discord.ext import commands

client = commands.Bot(command_prefix = '.')

@client.event
async def on_ready():
    print('Bot has connected to Discord!')

@client.command(aliases = ['rltracker', 'mmr'])
async def getStats(ctx, platform, player):

    base = 'https://rlstats.net/profile/'
    await ctx.send(base + platform + '/' + player)
    # response = requests.get(base + platform + '/' + player)
    # response = response.json()

    # data = response["main"]
    # currentTemperature = data["temp"]
    # weather = response["weather"]
    # description = weather[0]["description"]

@client.command()
@commands.has_permissions(manage_messages=True)
async def test(ctx):
    await ctx.send('You can manage messages.')

@client.command(aliases = ['cal', 'Cal', 'Calendar'])
@commands.has_permissions(external_emojis=True)
async def calendar(ctx):
    await ctx.send('You can use this command')
    cal1 = 'Practice Schedule: \nhttps://calendar.google.com/calendar/u/2?cid=Y19vbm82N2FvOTJoNjI3anFhZzZkcWRncnJlMEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t'
    cal2 = 'EGF Match Schedule: \nhttps://calendar.google.com/calendar/u/2?cid=Y19tMjJkbW9jc21vMGttcHRrOW1uNGgwZW10NEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t'
    cal3 = 'Team Meeting Schedule: \nhttps://calendar.google.com/calendar/u/2?cid=Y19qcTlyNWh1OGk1MTU1NWw2Mm82b2s1MGczZ0Bncm91cC5jYWxlbmRhci5nb29nbGUuY29t'
    await ctx.author.send('Here are the calendars:\n' + cal1 + cal2 + cal3)

@client.command(aliases = ['find_replays', 'replays'])
async def getReplays(ctx, platform, player):

    base = 'https://ballchasing.com/?'
    name = ''
    playlistString = ''
    season = ''
    minRank = ''
    title = ''
    title = ''
    title = ''
    title = ''
    title = ''
    await ctx.send(base +
                   'title=' + title + '&' +
                   'player-name=' + name + '&' +
                   playlistString +
                   'season=' + season + '&' +
                   'min-rank=' + minRank + '&' +
                   'max-rank=' + maxRank + '&' +
                   'map=' + map + '&' +
                   'replay-after=' + dateAfter + '&' +
                   'replay-before=' + dateBefore + '&' +
                   'upload-after=' + uploadBefore + '&' +
                   'upload-before=' + uploadAfter + '&')

client.run('Nzc5ODY0MTUzMTg3NjgwMjY3.X7mvFw.thfsHBohuOtcsYReJAX3nyvTXyQ')

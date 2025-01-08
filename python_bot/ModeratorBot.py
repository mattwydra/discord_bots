import discord
import certifi
import json
import requests
import urllib.parse
import random
from discord.ext import commands

client = commands.Bot(command_prefix = '!mod!')

@client.event
async def on_ready():
    print('Bot has connected to Discord!')

@commands.has_permissions(manage_messages=True)
@client.command(aliases = ['Clear'])
async def clear(ctx, amount = 5):
    await ctx.channel.purge(limit=amount)

@commands.has_permissions(manage_messages=True)
#@commands.has_role('RoleName')
@client.command(aliases = ['Testing'])
async def testing(ctx):
    await ctx.send('you may use this command')

@commands.has_permissions(manage_messages=True)
@client.command(aliases = ['Kick'])
async def kick(ctx, member : discord.Member, *, kickReason = None):
    await member.kick(reason = kickReason)

@commands.has_permissions(manage_messages=True)
@client.command(aliases = ['Ban'])
async def ban(ctx, member : discord.Member, *, banReason = None):
    await member.ban(reason = banReason)

@commands.has_permissions(manage_messages=True)
@client.command(aliases = ['UnBan', 'unBan', 'Unban'])
async def unban(ctx, *, member):
    banned_users = await ctx.guild.bans() #tuple (user, reason)
    member_name, member_discriminator = member.split('#')

    for ban_entry in banned_users:
        user = ban_entry.user #looping through all banned users

        if (user.name, user.discriminator) == (member_name, member_discriminator):
            await ctx.guild.unban(user)
            await ctx.send(f'Unbanned {user.mention}')
            return
# @client.command()
# async def clear(ctx, amount=0):
#     if (ctx.message.author.permissions_in(ctx.message.channel).manage_messages):
#         await ctx.channel.purge(limit= amount+1)
# @clear.error
# async def clear_error(ctx, error):
#     if isinstance(error, commands.MissingPermissions):
#         await ctx.send('Sorry you are not allowed to use this command.')
#
# @client.command()
# @commands.has_permissions(manage_messages=True)
# async def test(ctx):
#     await ctx.send('You can manage messages.')

client.run('TOKEN')

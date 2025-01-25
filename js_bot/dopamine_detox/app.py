import os

import discord
from dotenv import load_dotenv

load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')

import discord
from discord.ext import commands, tasks\

import asyncio
from datetime import datetime, timedelta, timezone
import pytz

# Set the timezone for EST
EST = pytz.timezone('US/Eastern')

print(TOKEN)


# Set up the bot
intents = discord.Intents.all()
intents.messages = True
bot = commands.Bot(command_prefix='!', intents=intents)

# Default schedule
schedule = [
    "No mindless entertainment (TV, social media, etc.).",
    "No processed food/sugars; follow the carnivore diet.",
    "Workout.",
    "Vision board (5 minutes).",
    "Emotional journaling (5 minutes).",
    "Affirmations.",
    "Read 10 pages of a book.",
    "Accomplish 3 major tasks.",
    "Plan tomorrow's top 3 actions."
]

# tasks = [
#     ""
# ]

@bot.event
async def on_ready():
    print(f'Logged in as {bot.user} (ID: {bot.user.id})')
    print('\nready\n\n------')

    # Start the task to send the message at 8:30 PM EST
    bot.loop.create_task(send_message_at_time())


async def send_message_at_time():
    # Get the current time in UTC
    now_utc = datetime.now(timezone.utc)

    # Get the target time (8:30 PM EST)
    target_time = EST.localize(datetime.combine(now_utc.date(), datetime.min.time())) + timedelta(hours=20, minutes=30)

    # If the target time has already passed today, schedule for tomorrow
    if now_utc > target_time.astimezone(timezone.utc):
        target_time += timedelta(days=1)

    # Calculate the time difference between now and the target time
    time_diff = target_time.astimezone(timezone.utc) - now_utc
    seconds_until_target = time_diff.total_seconds()

    # Wait until the target time
    await asyncio.sleep(seconds_until_target)

    # Send the message
    channel = bot.get_channel('1327802894975500360')  # Replace with your channel ID
    await channel.send("It's 8:30 PM EST! Time for your scheduled message!")

# Command to display the schedule
@bot.command()
async def test(ctx):
    send_message_at_time()

# Command to display the schedule
@bot.command()
async def daily(ctx):
    tasks = "\n".join(f"- {task}" for task in schedule)
    await ctx.send(f"**Daily Dopamine Detox Tasks:**\n{tasks}")

# Command to add a task
@bot.command()
async def add(ctx, *, task):
    if task not in schedule:
        schedule.append(task)
        await ctx.send(f"Task added: {task}")
    else:
        await ctx.send(f"Task already exists: {task}")

# Command to remove a task
@bot.command()
async def remove(ctx, *, task):
    if task in schedule:
        schedule.remove(task)
        await ctx.send(f"Task removed: {task}")
    else:
        await ctx.send(f"Task not found: {task}")

# Run the bot
bot.run(TOKEN)


# @client.command()
# async def clear(ctx, amount=0):
#     if (ctx.message.author.permissions_in(ctx.message.channel).manage_messages):
#         await ctx.channel.purge(limit= amount+1)
# @clear.error
# async def clear_error(ctx, error):
#     if isinstance(error, commands.MissingPermissions):
#         await ctx.send('Sorry you are not allowed to use this command.')
#

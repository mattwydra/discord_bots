import os

import discord
from dotenv import load_dotenv

load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')

import discord
from discord.ext import commands
import sqlite3

# Initialize bot
intents = discord.Intents.all()
intents.messages = True
bot = commands.Bot(command_prefix='!', intents=intents)

# Database setup
conn = sqlite3.connect('todo.db')
cursor = conn.cursor()
cursor.execute('''
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY,
        user_id TEXT NOT NULL,
        description TEXT NOT NULL,
        label TEXT
    )
''')
conn.commit()

# Add task
@bot.command()
async def add(ctx, description: str, label: str = None):
    cursor.execute('INSERT INTO tasks (user_id, description, label) VALUES (?, ?, ?)',
                   (str(ctx.author.id), description, label))
    conn.commit()
    await ctx.send(f'Task added: "{description}" with label "{label}"')

# Remove task
@bot.command()
async def remove(ctx, task_id: int):
    cursor.execute('DELETE FROM tasks WHERE id = ? AND user_id = ?', (task_id, str(ctx.author.id)))
    conn.commit()
    await ctx.send(f'Task with ID {task_id} removed.')

# View tasks
@bot.command()
async def view(ctx, label: str = None):
    if label:
        cursor.execute('SELECT id, description, label FROM tasks WHERE user_id = ? AND label = ?', 
                       (str(ctx.author.id), label))
    else:
        cursor.execute('SELECT id, description, label FROM tasks WHERE user_id = ?', 
                       (str(ctx.author.id),))
    tasks = cursor.fetchall()
    if tasks:
        response = '\n'.join([f'[{task[0]}] {task[1]} (Label: {task[2]})' for task in tasks])
    else:
        response = 'No tasks found.'
    await ctx.send(response)

# Run the bot
bot.run(TOKEN)

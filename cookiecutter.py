import os
from shutil import copyfile

import click
import git


def git_init(directory):
    g = git.cmd.Git(directory)
    g.init()

def gitignore(directory):
    file_gitingore = f'{directory}/.gitignore'
    if not os.path.isfile(file_gitingore):
        script_dir = os.path.dirname(os.path.realpath(__file__))
        copyfile(f'{script_dir}/.gitignore', file_gitingore)

def create_dir(directory):
    if not os.path.isdir(directory):
        os.mkdir(directory)
    if not os.path.isdir(f'{directory}/.git'):
        git_init(directory)

@click.command()
@click.option('--directory', default=os.getcwd(), help="project directory")
def main(directory):
    create_dir(directory)
    gitignore(directory)


if __name__ == "__main__":
    main()

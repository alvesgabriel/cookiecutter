from setuptools import setup


setup(
    name='cookiecutter',
    version='0.1',
    py_moduler=['cookiecutter'],
    install_requires=[
        'Click',
        'GitPython',
    ],
    entry_points='''
        [console_scripts]
        cookiecuter=cookiecutter:main
    ''',
)
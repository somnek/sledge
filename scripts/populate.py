import redis
import json
from rich.console import Console

c = Console()

r = redis.StrictRedis('localhost', 6666, db=0)
d = {
    'rob': 'go',
    'primeagen': ['rust', 'typescript', 'go'],
    'josh valty': {'lib': 'jax', 'model': ['efficientnet', 'resnet']},
    'codico': 'stare',
}


if __name__ == '__main__':
    if not r.get('j'):
        r.set('j', json.dumps(d))
        c.log(f'successfully inserted {d}')
    else:
        c.log('skip insert')
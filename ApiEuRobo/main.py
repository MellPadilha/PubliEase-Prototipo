from flask import Flask, request, jsonify
import classificador as cl

app = Flask(__name__)


@app.route('/motivo', methods=['POST'])
def receber_frase():
    data = request.get_json()
    if not data or 'frase' not in data:
        return jsonify({'error': 'O campo "frase" é obrigatório!'}), 400

    frase = data['frase']
        
    if frase == "":
        return jsonify({'message': str("SemAcesso")})

    result = cl.classificador_motivo(frase)

    return jsonify({'message': str(result)})

if __name__ == '__main__':
    app.run(debug=True, port=7024)

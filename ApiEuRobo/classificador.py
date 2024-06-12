import pickle
from sklearn.feature_extraction.text import CountVectorizer

target_names = {0: 'False', 1: 'True', 2: 'NaoMotivo'}

with open('model.pkl', 'rb') as model_file:
    loaded_model = pickle.load(model_file)

with open('vectorizer.pkl', 'rb') as vectorizer_file:
    loaded_vectorizer = pickle.load(vectorizer_file)
    
def predict_new_reasons(reasons):
    X_new = loaded_vectorizer.transform(reasons)
    y_new_pred = loaded_model.predict(X_new)
    y_new_pred_mapped = [target_names[val] for val in y_new_pred]
    return y_new_pred_mapped

def classificador_motivo(frase):
    reasons = [str(frase)]

    predictions = predict_new_reasons(reasons)
    for string, prediction in zip(reasons, predictions):
        return f"{prediction}"

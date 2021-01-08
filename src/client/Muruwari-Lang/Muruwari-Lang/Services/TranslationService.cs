using System.Net.Http;
using System.Threading.Tasks;
using Muruwari_Lang.Interfaces;
using Muruwari_Lang.Models;
using Newtonsoft.Json;

namespace Muruwari_Lang.Services
{
    public class TranslationService : ITranslationService
    {
        private IHttpClientFactory _client;

        public TranslationService(IHttpClientFactory httpClientFactory)
        {
            _client = httpClientFactory; 
        }

        public async Task<TranslationModel> GetTranslationAsync(string phrase)
        {
            try
            {
                var word = new TranslationInput();
                word.WORD = phrase;

                var content = new StringContent(JsonConvert.SerializeObject(word));

                var client = _client.CreateClient();
                var res = await client.PostAsync($"http://localhost:8000/search-phrase", content);

                if(res.IsSuccessStatusCode)
                {
                    var json = JsonConvert.DeserializeObject<TranslationModel>(await res.Content.ReadAsStringAsync());
                    return json;
                }

                return null;
            }
            catch
            {
                throw; 
            }
        }
    }
}

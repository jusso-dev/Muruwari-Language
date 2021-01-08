using System.Threading.Tasks;
using Muruwari_Lang.Models;

namespace Muruwari_Lang.Interfaces
{
    public interface ITranslationService
    {
        Task<TranslationModel> GetTranslationAsync(string phrase);
    }
}
